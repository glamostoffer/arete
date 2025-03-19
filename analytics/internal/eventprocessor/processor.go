package eventprocessor

import (
	"context"
	"sync"

	"github.com/IBM/sarama"
	"github.com/glamostoffer/arete/analytics/internal/domain"
	"github.com/gofiber/fiber/v2/log"
)

type processor struct {
	cfg      Config
	consumer sarama.PartitionConsumer
	repo     repository
	doneCh   chan (struct{})
}

func New(cfg Config, consumer sarama.PartitionConsumer, repo repository) *processor {
	doneCh := make(chan struct{}, 1)

	return &processor{
		cfg:      cfg,
		consumer: consumer,
		repo:     repo,
		doneCh:   doneCh,
	}
}

func (p *processor) Start(ctx context.Context) error {
	go func() {
		for {
			select {
			case err := <-p.consumer.Errors():
				log.Error(err)
			case msg := <-p.consumer.Messages():
				err := p.SaveEvent(ctx, msg)
				if err != nil {
					log.Error(err)
				}
			case <-p.doneCh:
				return
			}
		}
	}()

	return nil
}

func (p *processor) Stop(ctx context.Context) error {
	p.doneCh <- struct{}{}
	err := p.consumer.Close()
	return err
}

func (p *processor) SaveEvent(ctx context.Context, msg *sarama.ConsumerMessage) error {
	event := domain.Event{}
	err := event.ParseKafkaMessage(msg)
	if err != nil {
		return err
	}

	err = p.repo.InsertEvent(ctx, event)
	if err != nil {
		return err
	}

	return nil
}

func (p *processor) ProcessEvent(ctx context.Context) error {
	events, err := p.repo.SelectUnprocessedEvents(ctx, p.cfg.WorkersCount)
	if err != nil {
		return err
	}

	var wg sync.WaitGroup
	wg.Add(len(events))

	for _, event := range events {
		go func(event domain.Event) {
			defer wg.Done()
			var err error

			switch event.Type {
			case domain.QuizComplete:
				err = p.handleQuizEvent(event.Data, event.UserID, event.CourseID)
			case domain.TaskComplete:
				err = p.handleTaskEvent(event.Data, event.UserID, event.CourseID)
			}

			if err != nil {
				log.Error(err)
				return
			}

			err = p.repo.MarkProcessedEvent(ctx, event)
			if err != nil {
				log.Error(err)
				return
			}
		}(event)
	}

	wg.Wait()

	return nil
}
