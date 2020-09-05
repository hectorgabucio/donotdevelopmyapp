import React from 'react';
import ReactDOM from 'react-dom';
import './index.css';
import App from './App';
import * as serviceWorker from './serviceWorker';
import axios from 'axios';
import swal from 'sweetalert';

axios.defaults.baseURL = '/api';
axios.defaults.withCredentials = true;

// Add a 401 response interceptor
axios.interceptors.response.use(
  function (response) {
    return response;
  },
  function (error) {
    if (401 === error.response.status) {
      swal({
        title: 'Session Expired',
        text:
          'Your session has expired. You are going to be redirected to the login page',
        icon: 'warning',
        button: 'OK',
      })
        .then((value) => (window.location = '/auth'))
        .catch((error) => console.log(error));
    } else {
      return Promise.reject(error);
    }
  }
);

ReactDOM.render(
  <React.StrictMode>
    <App />
  </React.StrictMode>,
  document.getElementById('root')
);

// If you want your app to work offline and load faster, you can change
// unregister() to register() below. Note this comes with some pitfalls.
// Learn more about service workers: https://bit.ly/CRA-PWA
serviceWorker.unregister();
