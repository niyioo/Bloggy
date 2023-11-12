import React, { useState } from 'react';

const BlogPostForm = ({ onSubmit }) => {
  const [title, setTitle] = useState('');
  const [content, setContent] = useState('');

  const handleSubmit = (e) => {
    e.preventDefault();
    // Call the onSubmit prop with the form data
    onSubmit({ title, content });
  };

  return (
    <form onSubmit={handleSubmit}>
      <label>Title:</label>
      <input type="text" value={title} onChange={(e) => setTitle(e.target.value)} />

      <label>Content:</label>
      <textarea value={content} onChange={(e) => setContent(e.target.value)} />

      <button type="submit">Create Post</button>
    </form>
  );
};

export default BlogPostForm;
