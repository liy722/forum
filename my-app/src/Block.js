import React from 'react';
import { Button, Container, TextField } from '@mui/material';
import { Link as RouterLink } from 'react-router-dom';
import axios from 'axios';
import service from "./axios.js";
import {useState} from 'react';
import { useNavigate } from 'react-router-dom';
const Block = () => {
	const [count, setCount] = useState(0);
	const navigate = useNavigate();
	const handleSearch = (event) => {
	  
	  service.get('createtopic?topic='+document.getElementById('topic').value) // 替换为你的Spring Boot接口URL
	     .then(response => { 
			 setCount(count + 1);	
	         alert('success create topic!');
	  	     navigate('/blog');
	     })  
	     .catch(error => {  
	       console.log(error);  
	     });
		 
	};
	
  return (
    <Container style={{ display: 'flex', flexDirection: 'column', alignItems: 'center', justifyContent: 'center',
        minHeight: '100vh', }}>
      <TextField 
	    name="topic" 
		id="topic"
        label="topic"
        variant="outlined"
        style={{ width: '900px', minHeight: '150px', marginBottom: '16px' }}
      />
      <Button
        variant="contained"
        color="secondary"
        style={{ width: '200px' }}
        component={RouterLink} 
		onClick={handleSearch}
       
      >
        submit
      </Button>
    </Container>
  );
};

export default Block;
