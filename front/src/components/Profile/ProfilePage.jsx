import React, { useState } from 'react';
import { createAvatar } from '@dicebear/avatars';
import * as style from '@dicebear/avatars-identicon-sprites';
import SidebarMenu from '../Common/SidebarMenu';
import './ProfilePage.css';

const ProfilePage = () => {
  const [isMenuOpen, setIsMenuOpen] = useState(false);

  const generateAvatar = (seed, size = 150) => {
    return createAvatar(style, {
      seed,
      size,
      radius: 50,
      backgroundColor: "#ffffff",
    });
  };

  const userData = {
    name: "glamostoffer",
    email: "glamostoffer@gmail.com",
    registrationDate: "12 апреля 2025",
    globalRating: 6,
    courses: [
      { id: 1, name: "Основы Python", progress: 25 },
      { id: 3, name: "Алгоритмы и структуры данных", progress: 0 },
    ],
    courseRatings: {
      1: 512,
      2: 689,
      3: 812,
    },
    stats: {
      completedLessons: 24,
      totalLessons: 42,
      certificates: 2,
      streak: 7
    }
  };

  const topUsers = [
    { name: "test", rating: 30 },
    { name: "b2", rating: 25 },
    { name: "c3", rating: 20 },
    { name: "d4", rating: 15 },
    { name: "e5", rating: 10 },
  ];

  return (
    <div className="profile-layout">
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

      <div className="main-content">
        <header className="page-header">
          <h1>Профиль пользователя</h1>
        </header>

        <div className="profile-container">
          <div className="profile-content">
            <div className="profile-sidebar">
              <div className="avatar-container">
                <div
                  className="profile-avatar"
                  dangerouslySetInnerHTML={{ __html: generateAvatar(userData.name) }}
                />
              </div>

              <div className="profile-info">
                <div className="profile-info-item">
                  <div className="profile-info-label">Имя пользователя</div>
                  <div className="profile-info-value">{userData.name}</div>
                </div>

                <div className="profile-info-item">
                  <div className="profile-info-label">Электронная почта</div>
                  <div className="profile-info-value">{userData.email}</div>
                </div>

                <div className="profile-info-item">
                  <div className="profile-info-label">Дата регистрации</div>
                  <div className="profile-info-value">{userData.registrationDate}</div>
                </div>
              </div>

              <div className="delete-account-button-wrapper">
                <button className="delete-account-button">Удалить аккаунт</button>
              </div>
            </div>

            <div className="profile-main">
              <div className="top-row">
                <div className="test-of-the-day">
                  <h2>Тест дня</h2>
                  <p className="test-description">
                    Пройдите ежедневный тест и укрепите навыки программирования.
                    <br></br>
                    <br></br>
                    Сегодняшняя тема: «Структуры данных».
                  </p>
                  <button className="start-test-button">Начать тест</button>
                </div>

                <div className="rating-container">
                  <h2 className="courses-title">Ваш рейтинг</h2>
                  <div className="rating-values">
                    <div className="rating-item">
                      <span className="rating-label">Мировой рейтинг:</span>
                      <span className="rating-value">{userData.globalRating}</span>
                    </div>
                  </div>

                  <div className="top-users">
                    <h3>Топ-5 пользователей</h3>
                    {topUsers.map((user, index) => (
                      <div key={index} className="top-user">
                        <div
                          className="top-user-avatar"
                          dangerouslySetInnerHTML={{ __html: generateAvatar(user.name, 40) }}
                        />
                        <div className="top-user-info">
                          <div className="top-user-name">{user.name}</div>
                          <div className="top-user-rating">{user.rating}</div>
                        </div>
                      </div>
                    ))}
                  </div>
                </div>
              </div>

              <div className="courses-container">
                <h2 className="courses-title">Мои курсы</h2>
                {userData.courses.map(course => (
                  <div key={course.id} className="course-item">
                    <div>
                      <div className="course-name">{course.name}</div>
                      <div className="course-rating">
                        Рейтинг в курсе: {userData.courseRatings[course.id]}
                      </div>
                    </div>
                    <div className="course-progress-wrapper">
                      <div className="course-progress">
                        <div
                          className="course-progress-fill"
                          style={{ width: `${course.progress}%` }}
                        ></div>
                      </div>
                      <div className="course-progress-value">{course.progress}%</div>
                    </div>
                  </div>
                ))}
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  );
};

export default ProfilePage;