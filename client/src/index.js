import React from 'react';
import ReactDOM from 'react-dom/client';
import './index.css';
import App from './App';
import reportWebVitals from './reportWebVitals';

// const client = createClient({
//   url: 'http://localhost:3000/graphql',
//   exchanges: [dedupExchange, cacheExchange, fetchExchange],
// });
const root = ReactDOM.createRoot(document.getElementById('root'));


root.render(
      <App />
);

// If you want to start measuring performance in your app, pass a function
// to log results (for example: reportWebVitals(console.log))
// or send to an analytics endpoint. Learn more: https://bit.ly/CRA-vitals
reportWebVitals();
