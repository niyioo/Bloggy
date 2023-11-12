import React, { useState } from 'react';
import axios from 'axios';
import { useRouter } from 'next/router';

function RegisterForm() {
  const [formData, setFormData] = useState({ email: '', password: '' });
  const [error, setError] = useState(null);
<<<<<<< HEAD
  const router = useRouter();

  const handleInputChange = (e) => {
    const { name, value } = e.target;
    setFormData({ ...formData, [name]: value });
  };
=======

  const router = useRouter();
>>>>>>> 53834e38c4002610b63a29213c90384dab0149f8

  const handleSubmit = async (e) => {
    e.preventDefault();

    try {
<<<<<<< HEAD
      const response = await axios.post('http://localhost:8080/api/register', formData);
      router.push('/login');
    } catch (error) {
=======
      // Make an API request to your registration endpoint (replace with your actual URL)
      const response = await axios.post('/api/register', formData);

      // Handle successful registration, e.g., redirect to the login page
      router.push('/login');
    } catch (error) {
      // Handle registration errors, e.g., display an error message
>>>>>>> 53834e38c4002610b63a29213c90384dab0149f8
      setError('Registration failed. Please try again.');
    }
  };

  return (
    <div>
      <h2>Register</h2>
      {error && <div className="error-message">{error}</div>}
      <form onSubmit={handleSubmit}>
        <div>
          <label htmlFor="email">Email:</label>
          <input
            type="email"
            id="email"
            name="email"
            value={formData.email}
<<<<<<< HEAD
            onChange={handleInputChange}
=======
            onChange={(e) => setFormData({ ...formData, email: e.target.value })}
>>>>>>> 53834e38c4002610b63a29213c90384dab0149f8
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
<<<<<<< HEAD
            onChange={handleInputChange}
=======
            onChange={(e) => setFormData({ ...formData, password: e.target.value })}
>>>>>>> 53834e38c4002610b63a29213c90384dab0149f8
            required
          />
        </div>
        <div>
          <button type="submit">Register</button>
        </div>
      </form>
    </div>
  );
}

export default RegisterForm;
