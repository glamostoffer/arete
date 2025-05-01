import React from 'react';
import { Link } from 'react-router-dom';
import './SidebarMenu.css';

const SidebarMenu = ({ isOpen, onClose }) => {
  return (
    <div className={`sidebar ${!isOpen ? 'closed' : ''}`}>
      <div className="sidebar-header">
        <button className="toggle-menu" onClick={onClose}>
          {'✕'}
        </button>
      </div>

      <nav className="sidebar-nav">
        <div className="nav-items-container">
          <Link to="/courses" className="nav-item">
            <span className="nav-icon">{'</>'}</span>
            <span className="nav-text">Курсы</span>
          </Link>
          <Link to="/theory" className="nav-item">
            <span className="nav-icon">{'[]'}</span>
            <span className="nav-text">Теория</span>
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
      <div className="logout-button-wrapper">
        <button className="logout-button">Выйти</button>
      </div>
    </div>
  );
};

export default SidebarMenu;