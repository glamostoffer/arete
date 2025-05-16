package cache

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"github.com/glamostoffer/arete/practice/internal/domain"
	"github.com/gofrs/uuid"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

func (c *cache) getQuizzSessionKey(sessionID uuid.UUID) string {
	return fmt.Sprintf("quizzSession:%s", sessionID.String())
}

func (c *cache) SetQuizzSession(
	ctx context.Context,
	info domain.QuizzFinishedEvent,
	sessionID uuid.UUID,
	questionWithOpts map[domain.Question][]domain.QuestionOption,
	ttl time.Duration,
) error {
	data := make(map[string]interface{})

	infoData, err := json.Marshal(info)
	if err != nil {
		return err
	}

	result, err := json.Marshal(domain.QuizzSessionResult{
		QuestionsTotal:    0,
		RightAnswersCount: 0,
	})
	if err != nil {
		return err
	}

	data["info"] = infoData // todo make it const
	data["result"] = result // todo make it const

	for q, o := range questionWithOpts {
		questionWithOptions := domain.QuestionWithOpts{
			Question: q,
			Options:  o,
		}

		value, err := json.Marshal(questionWithOptions)
		if err != nil {
			return err
		}

		data[strconv.FormatInt(q.ID, 10)] = string(value)
	}

	pipe := c.cl.Pipeline()
	pipe.HSet(ctx, c.getQuizzSessionKey(sessionID), data)
	pipe.Expire(ctx, c.getQuizzSessionKey(sessionID), ttl)

	_, err = pipe.Exec(ctx)
	if err != nil {
		return err
	}

	return nil
}

func (c *cache) GetQuestionFromSession(
	ctx context.Context,
	sessionID uuid.UUID,
	questionID int64,
) (res domain.QuestionWithOpts, err error) {
	data, err := c.cl.HGet(ctx, c.getQuizzSessionKey(sessionID), strconv.FormatInt(questionID, 10)).Result()
	if err != nil {
		return res, err
	}

	err = json.Unmarshal([]byte(data), &res)
	if err != nil {
		return res, err
	}

	return res, nil
}

func (c *cache) GetDelQuestionFromSession(
	ctx context.Context,
	sessionID uuid.UUID,
	questionID int64,
) (res domain.QuestionWithOpts, err error) {
	data, err := c.cl.HGet(ctx, c.getQuizzSessionKey(sessionID), strconv.FormatInt(questionID, 10)).Result()
	if err != nil {
		return res, err
	}

	err = json.Unmarshal([]byte(data), &res)
	if err != nil {
		return res, err
	}

	if err = c.cl.HDel(ctx, c.getQuizzSessionKey(sessionID), strconv.FormatInt(questionID, 10)).Err(); err != nil {
		return res, err
	}

	return res, nil
}

func (c *cache) GetRandomQuestionFromSession(
	ctx context.Context,
	sessionID uuid.UUID,
) (res domain.QuestionWithOpts, err error) {
	data, err := c.cl.HGetAll(ctx, c.getQuizzSessionKey(sessionID)).Result()
	if err != nil {
		return res, err
	}

	var questionsWithOpts domain.QuestionWithOpts
	for k, v := range data {
		if k != "info" && k != "result" {
			err = json.Unmarshal([]byte(v), &questionsWithOpts)
			if err != nil {
				return res, err
			}

			break
		}
	}

	return questionsWithOpts, nil
}

func (c *cache) DelQuestionFromSession(
	ctx context.Context,
	sessionID uuid.UUID,
	questionID int64,
) error {
	if err := c.cl.HDel(ctx, c.getQuizzSessionKey(sessionID), strconv.FormatInt(questionID, 10)).Err(); err != nil {
		return err
	}

	return nil
}

func (c *cache) AddRightAnswerToQuizzSessionResult(
	ctx context.Context,
	sessionID uuid.UUID,
) error {
	res, err := c.cl.HGet(ctx, c.getQuizzSessionKey(sessionID), "result").Result() // todo make key const
	if err != nil {
		return err
	}

	cnt := gjson.GetBytes([]byte(res), "result.RightAnswersCount").Int()                 // todo make key const
	updatedResult, err := sjson.SetBytes([]byte(res), "result.RightAnswersCount", cnt+1) // todo make key const
	if err != nil {
		return err
	}

	err = c.cl.HSet(ctx, c.getQuizzSessionKey(sessionID), "result", string(updatedResult)).Err() // todo make key const
	if err != nil {
		return err
	}

	return nil
}

func (c *cache) GetInfoFromSession(
	ctx context.Context,
	sessionID uuid.UUID,
) (res domain.QuizzSessionResult, info domain.QuizzFinishedEvent, err error) {
	result, err := c.cl.HGet(ctx, c.getQuizzSessionKey(sessionID), "result").Result() // todo make key const
	if err != nil {
		return res, info, err
	}

	userQuizzIDsData, err := c.cl.HGet(ctx, c.getQuizzSessionKey(sessionID), "info").Result() // todo make key const
	if err != nil {
		return res, info, err
	}

	err = json.Unmarshal([]byte(result), &res)
	if err != nil {
		return res, info, err
	}

	err = json.Unmarshal([]byte(userQuizzIDsData), &info)
	if err != nil {
		return res, info, err
	}

	return res, info, nil
}
