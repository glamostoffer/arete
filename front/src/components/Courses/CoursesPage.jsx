import React, { useState } from 'react';
import SidebarMenu from '../Common/SidebarMenu';
import './CoursesPage.css';

const ClockIcon = () => (
  <svg width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" strokeWidth="2">
    <circle cx="12" cy="12" r="10"></circle>
    <polyline points="12 6 12 12 16 14"></polyline>
  </svg>
);

const DifficultyIcon = () => (
  <svg width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" strokeWidth="2">
    <path d="M10.29 3.86L1.82 18a2 2 0 0 0 1.71 3h16.94a2 2 0 0 0 1.71-3L13.71 3.86a2 2 0 0 0-3.42 0z"></path>
  </svg>
);

const CoursesPage = () => {
  const [isMenuOpen, setIsMenuOpen] = useState(false);
  const [activeFilter, setActiveFilter] = useState('all');
  const [enrolledCourses, setEnrolledCourses] = useState([1, 3]);

  const courses = [
    {
      id: 1,
      title: 'Программирование на Go (Golang)',
      description: 'Изучите язык программирования Go: многопоточность, эффективность и простота синтаксиса.',
      duration: '6 недель',
      difficulty: 'Средний',
      image: 'https://images.unsplash.com/photo-1617791160536-598cf32026fb?ixlib=rb-1.2.1&auto=format&fit=crop&w=600&q=80',
      category: 'programming'
    },
    {
      id: 2,
      title: 'Разработка на C#',
      description: 'Полный курс по C#: от основ до продвинутых концепций, включая .NET и разработку приложений.',
      duration: '7 недель',
      difficulty: 'Средний',
      image: 'https://images.unsplash.com/photo-1573164713988-8665fc963095?ixlib=rb-1.2.1&auto=format&fit=crop&w=600&q=80',
      category: 'programming'
    },
    {
      id: 3,
      title: 'Go для веб-разработки',
      description: 'Создание высоконагруженных веб-приложений с использованием Golang и современных фреймворков.',
      duration: '5 недель',
      difficulty: 'Продвинутый',
      image: 'https://images.unsplash.com/photo-1499951360447-b19be8fe80f5?ixlib=rb-1.2.1&auto=format&fit=crop&w=600&q=80',
      category: 'web'
    },
    {
      id: 4,
      title: 'C# и Unity для разработки игр',
      description: 'Основы разработки игр на Unity с использованием языка программирования C#.',
      duration: '8 недель',
      difficulty: 'Средний',
      image: 'https://images.unsplash.com/photo-1551103782-8ab07afd45c1?ixlib=rb-1.2.1&auto=format&fit=crop&w=600&q=80',
      category: 'gamedev'
    },
    {
      id: 5,
      title: 'Алгоритмы и структуры данных',
      description: 'Углубленное изучение алгоритмов и структур данных для подготовки к техническим собеседованиям.',
      duration: '6 недель',
      difficulty: 'Продвинутый',
      image: 'https://images.unsplash.com/photo-1558494949-ef010cbdcc31?ixlib=rb-1.2.1&auto=format&fit=crop&w=600&q=80',
      category: 'algorithms'
    },
    {
      id: 6,
      title: 'Основы баз данных',
      description: 'Изучение SQL и NoSQL баз данных, проектирование и оптимизация запросов.',
      duration: '5 недель',
      difficulty: 'Средний',
      image: 'https://images.unsplash.com/photo-1460925895917-afdab827c52f?ixlib=rb-1.2.1&auto=format&fit=crop&w=600&q=80',
      category: 'databases'
    },
    {
      id: 7,
      title: 'Мобильная разработка',
      description: 'Создание мобильных приложений для iOS и Android с использованием современных технологий.',
      duration: '8 недель',
      difficulty: 'Средний',
      image: 'https://images.unsplash.com/photo-1555774698-0b77e0d5fac6?ixlib=rb-1.2.1&auto=format&fit=crop&w=600&q=80',
      category: 'mobile'
    },
  ];

  const filters = [
    { id: 'all', name: 'Все курсы' },
    { id: 'programming', name: 'Программирование' },
    { id: 'web', name: 'Веб-разработка' },
    { id: 'algorithms', name: 'Алгоритмы' },
    { id: 'gamedev', name: 'Разработка игр' },
    { id: 'databases', name: 'Базы данных' },
    { id: 'mobile', name: 'Мобильная разработка' }
  ];

  const handleEnroll = (courseId) => {
    if (enrolledCourses.includes(courseId)) {
      setEnrolledCourses(enrolledCourses.filter(id => id !== courseId));
    } else {
      setEnrolledCourses([...enrolledCourses, courseId]);
    }
  };

  const filteredCourses = activeFilter === 'all' 
    ? courses 
    : courses.filter(course => course.category === activeFilter);

  return (
    <div className="courses-layout">
      <button 
        className="menu-toggle-button" 
        onClick={() => setIsMenuOpen(true)}
      >
        {'☰'}
      </button>

      <SidebarMenu 
        isOpen={isMenuOpen} 
        onClose={() => setIsMenuOpen(false)} 
      />

      <div className="courses-main-content">
        <header className="courses-header">
          <h1>Доступные курсы</h1>
          <p>Выберите курс, чтобы начать обучение</p>
        </header>

        <div className="filter-container">
          {filters.map(filter => (
            <button
              key={filter.id}
              className={`filter-button ${activeFilter === filter.id ? 'active' : ''}`}
              onClick={() => setActiveFilter(filter.id)}
            >
              {filter.name}
            </button>
          ))}
        </div>

        <div className="courses-grid">
          {filteredCourses.map(course => (
            <div key={course.id} className="course-card">
              <img src={course.image} alt={course.title} className="course-image" />
              <div className="course-content">
                <h3 className="course-title">{course.title}</h3>
                <p className="course-description">{course.description}</p>
                <div className="course-meta">
                  <span className="course-duration">
                    <ClockIcon /> {course.duration}
                  </span>
                  <span className="course-difficulty">
                    <DifficultyIcon /> {course.difficulty}
                  </span>
                </div>
                <div className="course-actions">
                  <button
                    className={`enroll-button ${enrolledCourses.includes(course.id) ? 'enrolled' : ''}`}
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
  );
};

export default CoursesPage;