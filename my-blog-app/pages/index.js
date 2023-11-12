import React from 'react';
import Link from 'next/link';

const styles = {
  container: {
    margin: 20,
    padding: 20,
    backgroundColor: '#f8f9fa', // Light background color
    color: '#343a40', // Dark text color
    fontFamily: 'Arial, sans-serif',
  },
  heading: {
    fontSize: '3rem',
    fontWeight: 'bold',
    textAlign: 'center',
    color: '#007bff', // Accent color
  },
  introText: {
    fontSize: '1.5rem',
    textAlign: 'center',
    marginBottom: '20px',
  },
  section: {
    margin: '20px 0',
    padding: '20px',
    border: '1px solid #ced4da', // Border color
    borderRadius: '8px', // Rounded corners
    backgroundColor: '#fff', // White background
  },
  featuredPost: {
    marginBottom: '15px',
  },
  latestPosts: {
    listStyle: 'none',
    padding: 0,
    margin: 0,
  },
  postListItem: {
    marginBottom: '10px',
  },
  buttonContainer: {
    display: 'flex',
    justifyContent: 'flex-end',
    marginTop: '20px',
  },
  button: {
    fontSize: '1.2rem',
    margin: '5px',
    padding: '10px 20px',
    backgroundColor: '#007bff', // Button background color
    color: '#fff', // Button text color
    borderRadius: '5px', // Rounded corners
    cursor: 'pointer',
  },
};

const Home = () => {
  return (
    <div style={styles.container}>
      <h1 style={styles.heading}>Bloggy</h1>
      <p style={styles.introText}>Discover and share your thoughts with the world.</p>

      <div style={styles.section}>
        <h2>Featured Blog Posts</h2>
        <div style={styles.featuredPost}>
          <p>Sample featured post 1</p>
        </div>
        <div style={styles.featuredPost}>
          <p>Sample featured post 2</p>
        </div>
      </div>

      <div style={styles.section}>
        <h2>Latest Blog Posts</h2>
        <ul style={styles.latestPosts}>
          <li style={styles.postListItem}>Sample latest post 1</li>
          <li style={styles.postListItem}>Sample latest post 2</li>
          <li style={styles.postListItem}>Sample latest post 3</li>
        </ul>
      </div>

      <div style={styles.buttonContainer}>
        <Link href="/register">
          <div style={styles.button}>Register</div>
        </Link>
        <span style={{ margin: '0 10px' }}>|</span>
        <Link href="/login">
          <div style={styles.button}>Login</div>
        </Link>
      </div>
    </div>
  );
};

export default Home;
