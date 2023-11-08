import React from 'react';
import Link from 'next/link';

const Home = () => {
  return (
    <div>
      <h1>Welcome to My Blog Platform</h1>
      <p>Discover and share your thoughts with the world.</p>

      <div>
        <h2>Featured Blog Posts</h2>
        {/* Display featured blog posts here */}
      </div>

      <div>
        <h2>Latest Blog Posts</h2>
        {/* Display a list of the latest blog posts here */}
      </div>

      <div>
        <Link href="/register">
          <a>Register</a>
        </Link>
        &nbsp;|&nbsp;
        <Link href="/login">
          <a>Login</a>
        </Link>
      </div>
    </div>
  );
};

export default Home;
