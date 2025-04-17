import React from 'react';
import { Link } from 'react-router-dom';
import './LoginPage.css';

const RegisterPage = () => {
  return (
    <div className="page-background">
      <div className="light-orb"></div>
      <div className="login-container">
        <div className="login-header">
          <div className="login-logo">Арете</div>
        </div>
        
        <div className="login-content">
          <h1 className="login-title">Создайте аккаунт</h1>
          
          <form className="login-form">
            <div className="form-group">
              <label htmlFor="email" className="form-label">Электронная почта</label>
              <input 
                type="email" 
                id="email" 
                className="form-input" 
                placeholder="Введите вашу почту" 
              />
            </div>
            
            <div className="form-group">
              <label htmlFor="username" className="form-label">Имя пользователя</label>
              <input 
                type="text" 
                id="username" 
                className="form-input" 
                placeholder="Придумайте имя пользователя" 
              />
            </div>
            
            <div className="form-group">
              <label htmlFor="password" className="form-label">Пароль</label>
              <input 
                type="password" 
                id="password" 
                className="form-input" 
                placeholder="Придумайте пароль" 
              />
            </div>
            
            <div className="form-group">
              <label htmlFor="confirm-password" className="form-label">Подтвердите пароль</label>
              <input 
                type="password" 
                id="confirm-password" 
                className="form-input" 
                placeholder="Повторите пароль" 
              />
            </div>
            
            <button type="submit" className="login-button">Зарегистрироваться</button>
          </form>
          
          <div className="login-links">
            <Link to="/login" className="login-link">Уже есть аккаунт? Войти</Link>
          </div>
        </div>
        
        <div className="login-footer">
          <p className="login-footer-text">
            Регистрируясь, вы соглашаетесь с нашими условиями использования
          </p>
        </div>
      </div>
    </div>
  );
};

export default RegisterPage;