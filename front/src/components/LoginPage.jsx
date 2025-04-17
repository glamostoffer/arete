import React from 'react';
import { Link } from 'react-router-dom';
import './LoginPage.css';

const LoginPage = () => {
  return (
    <div className="page-background">
      <div className="light-orb"></div>
      <div className="login-container">
        <div className="login-header">
          <div className="login-logo">Арете</div>
        </div>
        
        <div className="login-content">
          <h1 className="login-title">Войдите в ваш аккаунт</h1>
          
          <form className="login-form">
            <div className="form-group">
              <label htmlFor="username" className="form-label">Почта или логин</label>
              <input 
                type="text" 
                id="username" 
                className="form-input" 
                placeholder="Введите почту или имя пользователя" 
              />
            </div>
            
            <div className="form-group">
              <label htmlFor="password" className="form-label">Пароль</label>
              <input 
                type="password" 
                id="password" 
                className="form-input" 
                placeholder="Введите ваш пароль" 
              />
            </div>
            
            <button type="submit" className="login-button">Войти</button>
          </form>
          
          <div className="login-links">
            <Link to="/forgot-password" className="login-link">Забыли пароль?</Link>
            <Link to="/register" className="login-link">Создать аккаунт</Link>
          </div>
        </div>
        
        <div className="login-footer">
          <p className="login-footer-text">
            Если у вас возникли вопросы, свяжитесь с нашей поддержкой
          </p>
        </div>
      </div>
    </div>
  );
};

export default LoginPage;