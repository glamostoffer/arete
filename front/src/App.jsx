import React from "react";
import { BrowserRouter as Router, Routes, Route } from "react-router-dom";
import LoginPage from "./components/Auth/LoginPage.jsx";
import RegisterPage from "./components/Auth/RegisterPage.jsx";
import ProfilePage from "./components/Profile/ProfilePage.jsx";
import EmailConfirmationPage from "./components/Auth/EmailConfirmationPage.jsx";
import CoursesPage from "./components/Courses/CoursesPage.jsx";
import QuizzesPage from "./components/Quizz/QuizzesPage.jsx";
import TasksPage from "./components/Task/TasksPage.jsx";

const App = () => {
  return (
    <Router>
      <Routes>
        <Route path="/" element={<LoginPage />} />
        <Route path="/login" element={<LoginPage />} />
        <Route path="/register" element={<RegisterPage />} />
        <Route path="/profile" element={<ProfilePage />} />
        <Route path="/confirm-email" element={<EmailConfirmationPage />} />
        <Route path="/courses" element={<CoursesPage />} />
        <Route path="/quizzes" element={<QuizzesPage />} />
        <Route path="/Tasks" element={<TasksPage />} />
      </Routes>
    </Router>
  );
};

export default App;