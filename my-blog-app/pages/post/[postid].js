import React from 'react';
import { useRouter } from 'next/router';

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
};

const fetchPost = async (postId) => {
  try {
    console.log('Fetching post:', postId);
    const response = await fetch(`http://localhost:8080/api/blogposts/${postId}`);
    
    if (!response.ok) {
      console.error('Failed to fetch the individual post - HTTP error:', response.status);
      const errorMessage = await response.text();
      console.error('Error message:', errorMessage);
      throw new Error('Failed to fetch the individual post');
    }

    const post = await response.json();
    console.log('Fetched post:', post);
    return post;
  } catch (error) {
    console.error('Error fetching the individual post:', error);
    throw error;
  }
};

const Post = ({ post }) => {
  return (
    <div style={styles.container}>
      <h1 style={styles.heading}>{post.title}</h1>
      <p>{post.content}</p>
      <p>Author: {post.author}</p>
      <p>Created At: {post.created_at}</p>
    </div>
  );
};

export const getServerSideProps = async ({ params }) => {
  const { postId } = params;

  try {
    const post = await fetchPost(postId);
    
    return {
      props: {
        post,
      },
    };
  } catch (error) {
    console.error('Error fetching the individual post:', error);

    // Return a 404 page if the individual post is not found
    return {
      notFound: true,
    };
  }
};

export default Post;
