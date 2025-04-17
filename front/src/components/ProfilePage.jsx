import React, { useState } from 'react';
import { Link } from 'react-router-dom';
import { createAvatar } from '@dicebear/avatars';
import * as style from '@dicebear/avatars-identicon-sprites';
import './ProfilePage.css';

const ProfilePage = () => {
  const [isMenuOpen, setIsMenuOpen] = useState(true);
  
  const generateAvatar = (email) => {
    return createAvatar(style, {
      seed: email,
      size: 150,
      radius: 50,
      backgroundColor: "#ffffff",
    });
  };
  
  const userData = {
    name: "kuro",
    email: "kuroyuki25@gmail.com",
    registrationDate: "12 апреля 2025",
    overallProgress: 65,
    courses: [
      { id: 1, name: "Основы Python", progress: 80 },
      { id: 2, name: "Веб-разработка", progress: 45 },
      { id: 3, name: "Алгоритмы и структуры данных", progress: 30 },
    ],
    stats: {
      completedLessons: 24,
      totalLessons: 42,
      certificates: 2,
      streak: 7
    }
  };

  return (
    <div className="profile-layout">
    {/* Sidebar Menu */}
    {!isMenuOpen && (
      <div className="menu-toggle-button">
        <button className="toggle-menu" onClick={() => setIsMenuOpen(true)}>
          {'☰'}
        </button>
      </div>
    )}
    
    {isMenuOpen && (
      <div className="sidebar">
        <div className="sidebar-header">
          <button className="toggle-menu" onClick={() => setIsMenuOpen(false)}>
            {'✕'}
          </button>
        </div>
        
        <nav className="sidebar-nav">
          <div className="nav-items-container">
            <Link to="/courses" className="nav-item">
              <span className="nav-icon">{'</>'}</span>
              <span className="nav-text">Курсы</span>
            </Link>
            <Link to="/tasks" className="nav-item">
              <span className="nav-icon">{'{}'}</span>
              <span className="nav-text">Задачи</span>
            </Link>
            <Link to="/quizzes" className="nav-item">
              <span className="nav-icon">{'()'}</span>
              <span className="nav-text">Тесты</span>
            </Link>
            <Link to="/profile" className="nav-item">
              <span className="nav-icon">{'<>'}</span>
              <span className="nav-text">Профиль</span>
            </Link>
          </div>
        </nav>
      </div>
    )}

      {/* Main Content */}
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
                  dangerouslySetInnerHTML={{ __html: generateAvatar(userData.email) }}
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
              
              <div className="profile-actions">
                <Link to="/edit-profile" className="profile-button">Редактировать</Link>
                <button className="profile-button">Настройки</button>
              </div>
            </div>
            
            <div className="profile-main">
              <div className="progress-container">
                <h2 className="progress-title">Общий прогресс</h2>
                <div className="progress-bar">
                  <div 
                    className="progress-fill" 
                    style={{ width: `${userData.overallProgress}%` }}
                  ></div>
                </div>
                
                <div className="progress-stats">
                  <div className="progress-stat">
                    <div className="progress-stat-value">{userData.stats.completedLessons}</div>
                    <div className="progress-stat-label">Уроков пройдено</div>
                  </div>
                  
                  <div className="progress-stat">
                    <div className="progress-stat-value">{userData.stats.totalLessons}</div>
                    <div className="progress-stat-label">Всего уроков</div>
                  </div>
                  
                  <div className="progress-stat">
                    <div className="progress-stat-value">{userData.stats.certificates}</div>
                    <div className="progress-stat-label">Сертификатов</div>
                  </div>
                  
                  <div className="progress-stat">
                    <div className="progress-stat-value">{userData.stats.streak}</div>
                    <div className="progress-stat-label">Дней подряд</div>
                  </div>
                </div>
              </div>
              
              <div className="courses-container">
                <h2 className="courses-title">Мои курсы</h2>
                
                {userData.courses.map(course => (
                  <div key={course.id} className="course-item">
                    <div className="course-name">{course.name}</div>
                    <div className="course-progress">
                      <div 
                        className="course-progress-fill" 
                        style={{ width: `${course.progress}%` }}
                      ></div>
                    </div>
                    <div className="course-progress-value">{course.progress}%</div>
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