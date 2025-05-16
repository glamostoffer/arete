import React, { useEffect, useState } from 'react'
import SidebarMenu from '../Common/SidebarMenu'
import './CoursesPage.css'
import {
  useLearnCategoryStore,
  useLearnCourseStore,
  useLearnLessonStore,
} from '../../model/store'
import { useShallow } from 'zustand/react/shallow'

const ClockIcon = () => (
  <svg
    width='16'
    height='16'
    viewBox='0 0 24 24'
    fill='none'
    stroke='currentColor'
    strokeWidth='2'
  >
    <circle cx='12' cy='12' r='10'></circle>
    <polyline points='12 6 12 12 16 14'></polyline>
  </svg>
)

const DifficultyIcon = () => (
  <svg
    width='16'
    height='16'
    viewBox='0 0 24 24'
    fill='none'
    stroke='currentColor'
    strokeWidth='2'
  >
    <path d='M10.29 3.86L1.82 18a2 2 0 0 0 1.71 3h16.94a2 2 0 0 0 1.71-3L13.71 3.86a2 2 0 0 0-3.42 0z'></path>
  </svg>
)

const filters = [
  { id: undefined, name: 'Все курсы' },
  { id: 'programming', name: 'Программирование' },
  { id: 'web', name: 'Веб-разработка' },
  { id: 'algorithms', name: 'Алгоритмы' },
  { id: 'gamedev', name: 'Разработка игр' },
  { id: 'databases', name: 'Базы данных' },
  { id: 'mobile', name: 'Мобильная разработка' },
]

const CoursesPage = () => {
  const [isMenuOpen, setIsMenuOpen] = useState(false)
  const [activeFilter, setActiveFilter] = useState()
  const [enrolledCourses, setEnrolledCourses] = useState([1, 3])

  const { categories, getCategories } = useLearnCategoryStore(
    useShallow((state) => state)
  )

  const { courses, getCourses } = useLearnCourseStore(useShallow((state) => state))

  const { lessons, getLessons } = useLearnLessonStore(useShallow((state) => state))

  useEffect(() => {
    getCategories()
    getLessons(1)
  }, [])

  useEffect(() => {
    getCourses(activeFilter)
  }, [activeFilter])

  const handleEnroll = (courseId) => {
    if (enrolledCourses.includes(courseId)) {
      setEnrolledCourses(enrolledCourses.filter((id) => id !== courseId))
    } else {
      setEnrolledCourses([...enrolledCourses, courseId])
    }
  }

  const filteredCourses =
    activeFilter === undefined
      ? courses
      : courses.filter((course) => course.category === activeFilter)

  return (
    <div className='courses-layout'>
      <button className='menu-toggle-button' onClick={() => setIsMenuOpen(true)}>
        {'☰'}
      </button>

      <SidebarMenu isOpen={isMenuOpen} onClose={() => setIsMenuOpen(false)} />

      <div className='courses-main-content'>
        <header className='courses-header'>
          <h1>Доступные курсы</h1>
          <p>Выберите курс, чтобы начать обучение</p>
        </header>

        <div className='filter-container'>
          {filters
            .filter((fil) => !fil.id || categories.includes(fil.id))
            .map((filter) => (
              <button
                key={filter.id}
                className={`filter-button ${activeFilter === filter.id ? 'active' : ''}`}
                onClick={() => setActiveFilter(filter.id)}
              >
                {filter.name}
              </button>
            ))}
        </div>

        <div className='courses-grid'>
          {filteredCourses.map((course) => (
            <div key={course.id} className='course-card'>
              <img src={course.imageURL} alt={course.title} className='course-image' />
              <div className='course-content'>
                <h3 className='course-title'>{course.title}</h3>
                <p className='course-description'>{course.description}</p>
                <div className='course-meta'>
                  <span className='course-duration'>
                    <ClockIcon /> {course.duration}
                  </span>
                  <span className='course-difficulty'>
                    <DifficultyIcon /> {course.difficulty}
                  </span>
                </div>
                <div className='course-actions'>
                  <button
                    className={`enroll-button ${
                      enrolledCourses.includes(course.id) ? 'enrolled' : ''
                    }`}
                    onClick={() => handleEnroll(course.id)}
                  >
                    {enrolledCourses.includes(course.id) ? 'Записан' : 'Записаться'}
                  </button>
                </div>
              </div>
            </div>
          ))}
        </div>
      </div>
    </div>
  )
}

export default CoursesPage
