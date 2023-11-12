import React, { useState } from 'react';
import axios from 'axios';
import { useRouter } from 'next/router';

const LoginForm = () => {
  const [formData, setFormData] = useState({
    email: '',
    password: '',
  });
  const [error, setError] = useState(null);
<<<<<<< HEAD
=======

>>>>>>> 53834e38c4002610b63a29213c90384dab0149f8
  const router = useRouter();

  const handleInputChange = (e) => {
    const { name, value } = e.target;
    setFormData({ ...formData, [name]: value });
  };

  const isFormValid = () => {
    if (!formData.email || !formData.password) {
<<<<<<< HEAD
      setError('Please fill in all fields.');
=======
      setError("Please fill in all fields.");
>>>>>>> 53834e38c4002610b63a29213c90384dab0149f8
      return false;
    }

    if (!isValidEmail(formData.email)) {
<<<<<<< HEAD
      setError('Invalid email address.');
=======
      setError("Invalid email address.");
>>>>>>> 53834e38c4002610b63a29213c90384dab0149f8
      return false;
    }

    return true;
  };

  const isValidEmail = (email) => {
<<<<<<< HEAD
    return /^[A-Z0-9._%+-]+@[A-Z0-9.-]+\.[A-Z]{2,4}$/i.test(email);
=======
    return /^[\w-]+(\.[\w-]+)*@[\w-]+(\.[\w-]+)+$/.test(email);
>>>>>>> 53834e38c4002610b63a29213c90384dab0149f8
  };

  const handleSubmit = async (e) => {
    e.preventDefault();

    if (!isFormValid()) {
      return;
    }

    try {
      const response = await axios.post('/api/login', formData);
<<<<<<< HEAD
      const token = response.data.token;

      router.push('/dashboard');
    } catch (error) {
      setError('Login failed. Please check your credentials.');
=======

      // Handle successful login, e.g., store the token in a cookie or local storage
      const token = response.data.token;

      // Redirect the user to the dashboard or another page
      router.push('/dashboard');
    } catch (error) {
      setError("Login failed. Please check your credentials.");
>>>>>>> 53834e38c4002610b63a29213c90384dab0149f8
    }
  };

  return (
    <div>
      <h2>Login</h2>
      {error && <div className="error-message">{error}</div>}
      <form onSubmit={handleSubmit}>
        <div>
          <label htmlFor="email">Email:</label>
          <input
            type="email"
            id="email"
            name="email"
            value={formData.email}
            onChange={handleInputChange}
            required
          />
        </div>
        <div>
          <label htmlFor="password">Password:</label>
          <input
            type="password"
            id="password"
            name="password"
            value={formData.password}
            onChange={handleInputChange}
            required
          />
        </div>
        <div>
          <button type="submit">Login</button>
        </div>
      </form>
    </div>
  );
};

export default LoginForm;
