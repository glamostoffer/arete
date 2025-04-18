import React, { useState } from 'react';
import { Link } from 'react-router-dom';
import './LoginPage.css';

const LoginPage = () => {
  const [username, setUsername] = useState('');
  const [password, setPassword] = useState('');
  const isFormValid = username && password;

  const handleSubmit = (e) => {
    e.preventDefault();
    if (!isFormValid) return;

    // TODO: Реальный логин
    console.log('Вход:', { username, password });
  };

  return (
    <div className="page-background">
      <div className="light-orb"></div>
      <div className="login-container">
        <div className="login-header">
          <div className="login-logo">Арете</div>
        </div>
        
        <div className="login-content">
          <h1 className="login-title">Войдите в ваш аккаунт</h1>
          
          <form className="login-form" onSubmit={handleSubmit}>
            <div className="form-group">
              <label htmlFor="username" className="form-label">Почта или логин</label>
              <input 
                type="text" 
                id="username" 
                className="form-input" 
                placeholder="Введите почту или имя пользователя" 
                value={username}
                onChange={(e) => setUsername(e.target.value)}
              />
            </div>
            
            <div className="form-group">
              <label htmlFor="password" className="form-label">Пароль</label>
              <input 
                type="password" 
                id="password" 
                className="form-input" 
                placeholder="Введите ваш пароль" 
                value={password}
                onChange={(e) => setPassword(e.target.value)}
              />
            </div>
            
            <button
              type="submit"
              className={`login-button ${!isFormValid ? 'login-button-disabled' : ''}`}
              disabled={!isFormValid}
            >
              Войти
            </button>
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
