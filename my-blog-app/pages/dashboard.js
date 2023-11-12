import React from 'react';

const styles = {
  container: {
    margin: '20px',
    padding: '20px',
    maxWidth: '800px',
    border: '1px solid #ccc',
    borderRadius: '8px',
    backgroundColor: '#fff',
    boxShadow: '0 0 10px rgba(0, 0, 0, 0.1)',
  },
  heading: {
    fontSize: '24px',
    fontWeight: 'bold',
    marginBottom: '20px',
    textAlign: 'center',
  },
  content: {
    display: 'grid',
    gridTemplateColumns: 'repeat(auto-fill, minmax(300px, 1fr))',
    gap: '20px',
  },
  card: {
    padding: '20px',
    border: '1px solid #ccc',
    borderRadius: '8px',
    backgroundColor: '#f4f4f4',
    textAlign: 'center',
  },
  cardTitle: {
    fontSize: '20px',
    fontWeight: 'bold',
    marginBottom: '10px',
  },
  cardButton: {
    padding: '10px',
    backgroundColor: '#007bff',
    color: '#fff',
    border: 'none',
    borderRadius: '5px',
    cursor: 'pointer',
  },
};

const Dashboard = () => {
  return (
    <div style={styles.container}>
      <h1 style={styles.heading}>Welcome to Your Blog Dashboard</h1>
      <div style={styles.content}>
        <div style={styles.card}>
          <h2 style={styles.cardTitle}>My Blog Posts</h2>
          <p>You can manage and view your blog posts here.</p>
          <button style={styles.cardButton}>View Posts</button>
        </div>
        <div style={styles.card}>
          <h2 style={styles.cardTitle}>Create a New Post</h2>
          <p>Start writing a new blog post and share your thoughts.</p>
          <button style={styles.cardButton}>Create Post</button>
        </div>
      </div>
    </div>
  );
};

export default Dashboard;
