import React, { useEffect, useState } from 'react';  
import { TextField, Button, Container, Grid, Link } from '@mui/material';
import { Link as RouterLink } from 'react-router-dom';
import { Box } from '@mui/material';
import { Table, TableHead, TableBody, TableRow, TableCell, Paper } from '@mui/material';
import axios from 'axios';
import service from "./axios.js";


const Thread = () => {  
  const [data, setData] = useState([]);  
  const [data1, setData1] = useState(0);
  var i=0;
  
  const myDelete = (id) => {
        var info=document.getElementById(id);
		if(info!=null&&info!=undefined){
  		service.get('deletesubthread?stid='+info.innerText) 
  		   .then(response => {
  			   if(response.data.message=='delete reply success!'){
  	             alert("Delete Reply Success!");
  			     info.parentElement.parentElement.remove();
  			   }
  			   
  			})  
  		   .catch(error => {  
  		     console.log(error);  
  		   });
  	    			
  		}
    };
	
	
  useEffect(() => {  
    // Call axios here to fetch data
    const fetchData = async () => {  
      try {  

    i++;
	{/*thread?username=liy&topic=abc&tid=49*/}
    const urlParams = new URLSearchParams(window.location.search);
	const topic = urlParams.get('topic');   
	const tid = urlParams.get('tid'); 
	const username = urlParams.get('username'); 
	
	
	document.getElementById("topic").innerHTML="<b>"+topic+"</b>";
	document.getElementById("username").innerHTML="<b>"+username+"</b>";
    if(i==1)
	service.get('subthread?tid='+tid)
	   .then(response => {
		  setData1(data1+1);
          setData(response.data);
		 
	   })  
	   .catch(error => {  
	     console.log(error);  
	});


      } catch (error) {  
        console.error('Error fetching data:', error);  
      }  
    };  
  
  
  
    fetchData();  
  }, []); 
    return (
    <Container>
      <>    
        <Grid container>
          <Grid item xs={12}>
         { 
    	   <Table>
        
    	
        <TableBody>
    	<TableRow>
    	
    	<TableCell id="topic" style={{width: "80%"}}></TableCell>
    	<TableCell id="username" style={{width: "10%"}}></TableCell>
    	<TableCell style={{width: "10%"}}>&nbsp;</TableCell>
    	</TableRow>
	{/*	revert.revert+","+revert.rusername+","+revert.stid+","+mydelete*/}
    	 {data.map((item, index) => (  
          <TableRow key={index}>
            <TableCell>{item.slice(0, item.indexOf(","))}</TableCell>
            <TableCell>{item.split(",")[1]}</TableCell>
    		<TableCell>
		{ item.split(",")[3]==='1'&&
		<Link  onClick={() => myDelete(index)}>delete</Link>}  
		<div  id={index} style={{display: 'none'}}>{item.split(",")[2]}</div>
			</TableCell>
          </TableRow>
    	      ))}    
          {/* Add more rows as needed */}
        </TableBody>
      </Table>
      }
          </Grid>
    	  <Grid item xs={12}>
    	    {
    			<Button
          variant="contained"
          color="primary"
          style={{ marginTop: '8px' }}
    	  component={RouterLink}
    	  to="/reply"
        >
          Reply
        </Button>
		
    	}&nbsp;&nbsp;
		{<Button
		  variant="contained"
		  color="secondary"
		  style={{ marginTop: '8px' }}
		  component={RouterLink}
		  to="/blog"
		>
		  Back to dashboard
		</Button>}
    	  </Grid>
          
        </Grid>
      </>
    </Container>
  );  
};  
  
export default Thread;