// QuizzesPage.jsx
import React, { useState } from 'react';
import SidebarMenu from '../Common/SidebarMenu';
import './QuizzesPage.css';

const QuizzesPage = () => {
  const [isMenuOpen, setIsMenuOpen] = useState(false);
  const [selectedCourse, setSelectedCourse] = useState(null);
  const [currentQuiz, setCurrentQuiz] = useState(null);
  const [currentQuestionIndex, setCurrentQuestionIndex] = useState(0);
  const [selectedAnswer, setSelectedAnswer] = useState(null);
  const [showExplanation, setShowExplanation] = useState(false);
  const [completedQuizzes, setCompletedQuizzes] = useState([]);
  const [correctStreak, setCorrectStreak] = useState(0);
  const [isStreakActive, setIsStreakActive] = useState(false);
  const [streak, setStreak] = useState(0);

  const courses = [
    { id: 1, name: 'Основы Python' },
    { id: 2, name: 'Веб-разработка' },
    { id: 3, name: 'Алгоритмы и структуры данных' }
  ];

  const quizzes = {
    1: [
      {
        id: 101,
        title: 'Основы синтаксиса Python',
        description: 'Проверьте свои знания базового синтаксиса Python',
        questions: [
          {
            id: 1,
            text: 'Какой оператор используется для объявления функции в Python?',
            options: [
              { id: 1, text: 'function', isCorrect: false },
              { id: 2, text: 'def', isCorrect: true },
              { id: 3, text: 'func', isCorrect: false },
              { id: 4, text: 'declare', isCorrect: false }
            ],
            explanation: 'В Python функции объявляются с помощью ключевого слова "def". Например: `def my_function():`'
          },
          {
            id: 2,
            text: 'Какой тип данных является неизменяемым в Python?',
            options: [
              { id: 1, text: 'Список (list)', isCorrect: false },
              { id: 2, text: 'Словарь (dict)', isCorrect: false },
              { id: 3, text: 'Кортеж (tuple)', isCorrect: true },
              { id: 4, text: 'Множество (set)', isCorrect: false }
            ],
            explanation: 'Кортежи (tuples) являются неизменяемыми. После создания элементы кортежа нельзя изменить. Списки, словари и множества можно изменять.'
          },
          {
            id: 3,
            text: 'Как правильно создать список в Python?',
            options: [
              { id: 1, text: 'list = (1, 2, 3)', isCorrect: false },
              { id: 2, text: 'list = [1, 2, 3]', isCorrect: true },
              { id: 3, text: 'list = {1, 2, 3}', isCorrect: false },
              { id: 4, text: 'list = <1, 2, 3>', isCorrect: false }
            ],
            explanation: 'Списки создаются с использованием квадратных скобок. Круглые скобки создают кортеж, фигурные — множество или словарь.'
          },
          {
            id: 4,
            text: 'Что выведет этот код: `print(3 * "hi")`?',
            options: [
              { id: 1, text: 'hihihi', isCorrect: true },
              { id: 2, text: '3hi', isCorrect: false },
              { id: 3, text: 'hi hi hi', isCorrect: false },
              { id: 4, text: 'Ошибку', isCorrect: false }
            ],
            explanation: 'В Python умножение строки на число повторяет строку указанное количество раз. Это полезно для быстрого создания повторяющихся строк.'
          }
        ]
      },
      {
        id: 102,
        title: 'Условные операторы',
        description: 'Тест по условным операторам if-elif-else',
        locked: true,
        questions: []
      }
    ],
    2: [
      {
        id: 201,
        title: 'Основы HTML',
        description: 'Проверьте свои знания базового HTML',
        questions: []
      }
    ],
    3: [
      {
        id: 301,
        title: 'Сложность алгоритмов',
        description: 'Тест по оценке сложности алгоритмов',
        questions: []
      }
    ]
  };

  const handleStartQuiz = (quiz) => {
    setCurrentQuiz(quiz);
    setCurrentQuestionIndex(0);
    setSelectedAnswer(null);
    setShowExplanation(false);
  };

  const handleAnswerSelect = (answerId) => {
    setSelectedAnswer(answerId);
    const isCorrect = currentQuiz.questions[currentQuestionIndex].options.find(o => o.id === answerId)?.isCorrect;
    if (isCorrect) {
      setStreak(prev => prev + 1);
    } else {
      setStreak(0);
    }
    setShowExplanation(true);
  };
  

  const handleNextQuestion = () => {
    if (currentQuestionIndex < currentQuiz.questions.length - 1) {
      setCurrentQuestionIndex(currentQuestionIndex + 1);
      setSelectedAnswer(null);
      setShowExplanation(false);
    } else {
      handleCompleteQuiz();
    }
  };

  const handleCompleteQuiz = () => {
    if (currentQuiz) {
      setCompletedQuizzes([...completedQuizzes, currentQuiz.id]);
      setCurrentQuiz(null);
      setCurrentQuestionIndex(0);
      setCorrectStreak(0);
      setIsStreakActive(false);
    }
  };

  const isQuizUnlocked = (quiz) => {
    if (!quiz.locked) return true;
    const courseQuizzes = quizzes[selectedCourse];
    const quizIndex = courseQuizzes.findIndex(q => q.id === quiz.id);
    return quizIndex > 0 && completedQuizzes.includes(courseQuizzes[quizIndex - 1].id);
  };

  return (
    <div className="quizzes-layout">
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

      <div className="quizzes-main-content">
        <header className="quizzes-header">
          <h1>Тесты и квизы</h1>
          <p>Проверьте свои знания и закрепите изученный материал</p>
        </header>

        {!currentQuiz ? (
          <>
            <div className="course-selector-container">
              <select
                className="course-selector"
                value={selectedCourse || ''}
                onChange={(e) => setSelectedCourse(e.target.value ? parseInt(e.target.value) : null)}
              >
                <option value="">Выберите курс</option>
                {courses.map(course => (
                  <option key={course.id} value={course.id}>{course.name}</option>
                ))}
              </select>
            </div>

            {selectedCourse && (
              <div className="quizzes-container">
                <h2 className="quizzes-title">Тесты по курсу: {courses.find(c => c.id === selectedCourse).name}</h2>
                <div className="quizzes-grid">
                  {quizzes[selectedCourse]?.map(quiz => (
                    <div 
                      key={quiz.id} 
                      className={`quiz-card ${!isQuizUnlocked(quiz) ? 'locked' : ''} ${completedQuizzes.includes(quiz.id) ? 'completed' : ''}`}
                    >
                      <div className="quiz-content">
                        <div className="quiz-header">
                          <h3 className="quiz-title">{quiz.title}</h3>
                          {completedQuizzes.includes(quiz.id) && (
                            <span className="quiz-completed-badge">✓</span>
                          )}
                        </div>
                        <p className="quiz-description">{quiz.description}</p>
                        {!isQuizUnlocked(quiz) && (
                          <div className="quiz-lock-message">
                            Пройдите предыдущий тест для разблокировки
                          </div>
                        )}
                      </div>
                      <div className="quiz-actions">
                        <button
                          className={`start-quiz-button ${!isQuizUnlocked(quiz) ? 'disabled' : ''}`}
                          onClick={() => isQuizUnlocked(quiz) && handleStartQuiz(quiz)}
                          disabled={!isQuizUnlocked(quiz)}
                        >
                          {completedQuizzes.includes(quiz.id) ? 'Повторить тест' : 'Начать тест'}
                        </button>
                      </div>
                    </div>
                  ))}
                </div>
              </div>
            )}
          </>
        ) : (
          <div className="quiz-questions-container">
            <div className="quiz-progress-container">
              <div 
                className="quiz-progress-bar"
                style={{
                  width: `${((currentQuestionIndex + 1) / currentQuiz.questions.length) * 100}%`,
                  background: 'linear-gradient(to right, #00bcd4, #ffa726)'
                }}
              ></div>
              <div className="quiz-progress-text">
                Вопрос {currentQuestionIndex + 1} из {currentQuiz.questions.length}
              </div>
            </div>

            <div className="quiz-header">
              <h2>{currentQuiz.title}</h2>
              <p>{currentQuiz.description}</p>
            </div>

            <div className={`streak-counter ${streak > 0 ? 'active' : 'inactive'}`}>
                <span className="streak-fire">{streak > 0 ? '🔥' : '🕯️'}</span>
                Серия: {streak}
            </div>


            <div className="question-container">
              <h3 className="question-text">
                {currentQuiz.questions[currentQuestionIndex].text}
              </h3>
              
              <div className="options-container">
                {currentQuiz.questions[currentQuestionIndex].options.map(option => (
                  <div 
                    key={option.id}
                    className={`option ${selectedAnswer === option.id ? (option.isCorrect ? 'correct' : 'incorrect') : ''}`}
                    onClick={() => !showExplanation && handleAnswerSelect(option.id)}
                  >
                    {option.text}
                  </div>
                ))}
              </div>

              {showExplanation && (
                <div className="explanation-container">
                  <p className="explanation-text">
                    {selectedAnswer 
                      ? currentQuiz.questions[currentQuestionIndex].options.find(o => o.id === selectedAnswer).isCorrect
                        ? 'Правильно! 🎉'
                        : 'Неправильно 😕'
                      : ''}
                  </p>
                  <p className="explanation-text">
                    {currentQuiz.questions[currentQuestionIndex].explanation}
                  </p>
                  <button 
                    className="next-button"
                    onClick={handleNextQuestion}
                  >
                    {currentQuestionIndex < currentQuiz.questions.length - 1 
                      ? 'Следующий вопрос' 
                      : 'Завершить тест'}
                  </button>
                </div>
              )}
            </div>
          </div>
        )}
      </div>
    </div>
  );
};

export default QuizzesPage;