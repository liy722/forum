import React from 'react';
import { Button, Container, TextField } from '@mui/material';
import { Link as RouterLink } from 'react-router-dom';
import axios from 'axios';
import service from "./axios.js";
import { useNavigate } from 'react-router-dom';
const Reply = () => {
	const navigate = useNavigate();
  	const MySubmit = (event) => {
		service.get('reply?reply='+document.getElementById('reply').value) 
		  			
		  			 .then(response => {  
		  		  
		  		   alert('Reply Success!');
		  		   navigate('/thread'+response.data.message);
		  		   
		  		 })  
		  		 .catch(error => {  
		  		   console.log(error);  
		  		 });
		  	};
  return (
    <Container style={{ display: 'flex', flexDirection: 'column', alignItems: 'center', justifyContent: 'center',
        minHeight: '100vh', }}>
      <TextField
	    name="reply" 
		id="reply"
        label="reply"
        variant="outlined"
        style={{ width: '900px', minHeight: '150px', marginBottom: '16px' }}
      />
      <Button
        variant="contained"
        color="secondary"
        style={{ width: '200px' }}
        component={RouterLink}
        onClick={MySubmit}
      >
        submit
      </Button>
    </Container>
  );
};

export default Reply;
