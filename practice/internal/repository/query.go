package repository

const (
	querySelectQuizzesByCourseID = `
	select
		q.id,
		q.course_id,
		q.title,
		q."description",
		q.passing_score,
		q.sequence_number
		case 
			when q.sequence_number = 1 then false
			when exists (
				select 1 
				from public.quizz prev_q
				join public.user_completed_quizz prev_ucq on prev_q.id = prev_ucq.quizz_id
				where prev_q.course_id = q.course_id
				and prev_q.sequence_number = q.sequence_number - 1
				and prev_ucq.user_id = $2
			) then false
			else true
    	end AS is_locked
	from
		public.quizz q
	where 
		q.course_id = $1
	order by sequence_number
	limit $3
	offset $4;
	`
	queryGetQuizz = `
	select
		q.id,
		q.course_id,
		q.title,
		q."description",
		q.passing_score,
		q.sequence_number
		case 
			when q.sequence_number = 1 then false
			when exists (
				select 1 
				from public.quizz prev_q
				join public.user_completed_quizz prev_ucq on prev_q.id = prev_ucq.quizz_id
				where prev_q.course_id = q.course_id
				and prev_q.sequence_number = q.sequence_number - 1
				and prev_ucq.user_id = $2
			) then false
			else true
    	end AS is_locked
	from
		public.quizz q
	where 
		q.id = $1;
	`
	querySelectQuizzQuestions = `
	select
		id,
		quizz_id,
		question,
		explanation
	from
		public.question
	where
		quizz_id = $1;
	`
	querySelectQuestionOptions = `
	select
		id,
		question_id,
		"option",
		is_correct
	from 
		public.question_option
	where
		question_id = $1;
	`
	querySelectAllQuizzOptions = `
	select
		id,
		question_id,
		"option",
		is_correct
	from
		public.question_option
	where
		question_id = any($1);
	`
	queryMarkQuizzCompleted = `
	insert into public.user_completed_quizz(
		user_id,
		quizz_id
	) values (
		$1,
		$2
	);
	`
)
