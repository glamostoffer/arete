import React, { useState } from 'react';
import { Link, useNavigate } from 'react-router-dom';
import './LoginPage.css';

const RegisterPage = () => {
  const [email, setEmail] = useState('');
  const [username, setUsername] = useState('');
  const [password, setPassword] = useState('');
  const [confirmPassword, setConfirmPassword] = useState('');
  const [emailTouched, setEmailTouched] = useState(false);
  const [confirmTouched, setConfirmTouched] = useState(false);

  const navigate = useNavigate();

  const isEmailValid = /^[^\s@]+@[^\s@]+\.[^\s@]+$/.test(email);
  const doPasswordsMatch = password === confirmPassword;
  const isFormValid = email && isEmailValid && username && password && confirmPassword && doPasswordsMatch;

  const handleSubmit = (e) => {
    e.preventDefault();
    if (!isFormValid) return;

    // Здесь может быть отправка данных на сервер, потом:
    navigate('/confirm-email', { state: { email } }); // можно передать email как state
  };

  return (
    <div className="page-background">
      <div className="light-orb"></div>
      <div className="login-container">
        <div className="login-header">
          <div className="login-logo">Арете</div>
        </div>

        <div className="login-content">
          <h1 className="login-title">Создайте аккаунт</h1>

          <form className="login-form" onSubmit={handleSubmit}>
            <div className="form-group">
              <label htmlFor="email" className="form-label">Электронная почта</label>
              <input 
                type="email" 
                id="email" 
                className="form-input" 
                placeholder="Введите вашу почту" 
                value={email}
                onChange={(e) => setEmail(e.target.value)}
                onBlur={() => setEmailTouched(true)}
              />
              {emailTouched && !isEmailValid && (
                <div className="form-error">Введите корректный email</div>
              )}
            </div>

            <div className="form-group">
              <label htmlFor="username" className="form-label">Имя пользователя</label>
              <input 
                type="text" 
                id="username" 
                className="form-input" 
                placeholder="Придумайте имя пользователя" 
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
                placeholder="Придумайте пароль" 
                value={password}
                onChange={(e) => setPassword(e.target.value)}
              />
            </div>

            <div className="form-group">
              <label htmlFor="confirm-password" className="form-label">Подтвердите пароль</label>
              <input 
                type="password" 
                id="confirm-password" 
                className="form-input" 
                placeholder="Повторите пароль" 
                value={confirmPassword}
                onChange={(e) => setConfirmPassword(e.target.value)}
                onBlur={() => setConfirmTouched(true)}
              />
              {confirmTouched && !doPasswordsMatch && (
                <div className="form-error">Пароли не совпадают</div>
              )}
            </div>

            <button
              type="submit"
              className={`login-button ${!isFormValid ? 'login-button-disabled' : ''}`}
              disabled={!isFormValid}
            >
              Зарегистрироваться
            </button>
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
