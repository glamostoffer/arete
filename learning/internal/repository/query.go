package repository

const (
	queryGetCourses = `
	select
		c.id,
		c.title,
		c.description,
		c.duration,
		c.difficulty,
		cc.name as category,
		c.image_url,
		case when uc.user_id is not null then true else false end as is_enrolled
	from
		public.course c
		left join public.user_course uc on c.id = uc.course_id and uc.user_id = $1
		left join public.course_category cc on c.category_id = cc.id
	where
		$2::text[] is null or cc.name = any($2::text[])
	order by c.id
	limit ($3 + 1)
	offset $4;
	`

	queryGetCourseCategories = `
	select
		name
	from 
		public.course_category
	`
)

const (
	queryGetLessons = `
	select
		id,
		course_id,
		title,
		description,
		duration
	from
		public.lesson
	where
		course_id = $1
	order by id
	limit ($2 + 1)
	offset $3;
	`

	queryGetLessonDetails = `
	select
		l.id,
		l.course_id,
		l.title,
		l.description,
		l.duration,
		lc.content
	from 
		public.lesson l
		left join public.lesson_content lc on l.id = lc.lesson_id
	where
		l.id = $1;
	`
)
