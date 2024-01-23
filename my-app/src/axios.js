import axios from 'axios' 
 const service = axios.create({
   baseURL: 'http://localhost:8080/' // API Base URL
   
   
 });  
   
 export default service;
