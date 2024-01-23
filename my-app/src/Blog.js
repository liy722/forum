import React, { useState } from 'react';
import { TextField, Button, Container, Grid, Link } from '@mui/material';
import { Link as RouterLink } from 'react-router-dom';
import { Box } from '@mui/material';
import { Table, TableHead, TableBody, TableRow, TableCell, Paper } from '@mui/material';
import axios from 'axios';
import service from "./axios.js";
import { useNavigate } from 'react-router-dom';
const Blog = () => {
  const [searchQuery, setSearchQuery] = useState('');
  const navigate = useNavigate();
  const [data, setData] = useState([]);
  var i=0; 
  const handleSearch = (event) => {

	    if(i>0){
			window.location.reload();    //Refresh the page
		}
		service.get('blog?topic='+document.getElementById('topic').value)
		   .then(response => {
			   
		      var info= response.data;
			  setData(info);
			 
		   })  
		   .catch(error => {  
		     console.log(error);  
		   }); 
    
  };
  
  
  
 
  	
  const myDelete = (id) => { 
	    i++;
        var info0=document.getElementById("tr"+id);
		var info=document.getElementById(id);
		if(info!=null&&info!=undefined){
		service.get('deletetopic?tid='+info.innerText) 
		   .then(response => {
			   if(response.data.message=='delete topic success!'){
	             alert("Delete Topic Success!");
				 info0.innerHTML="";
			   }
			   
			})  
		   .catch(error => {  
		     console.log(error);  
		   });
	    			
		}
    };
	
  return (
    <Container>
      <>    
	  <Box display="flex">
      <div style={{ marginTop: '30px', marginLeft: '5px' }}>
	  <TextField id="topic"
	      name="topic"
          label="Search"
          variant="outlined"
          fullWidth
          value={searchQuery}
          onChange={(e) => setSearchQuery(e.target.value)}
          onKeyPress={(e) => {
            if (e.key === 'Enter') {
              handleSearch();
            }
          }}
        /></div>
      <div style={{ marginTop: '30px', marginLeft: '5px' }}>     
	  <Button
          variant="contained"
          color="primary"
          onClick={handleSearch}
          style={{ marginTop: '8px' }}
        >
          Search
        </Button></div>
      <div style={{ marginTop: '30px', marginLeft: '5px' }}> 
	  <Link component={RouterLink} to="/main" style={{ textDecoration: 'none' }}>
          <Button
            variant="contained"
            color="secondary"
            style={{ marginTop: '8px', marginLeft: '8px' }}
			component={RouterLink}
					  to="/block"
          >
            create my own thread
          </Button>
        </Link></div>
		
		
    </Box>
        <Grid container>
          <Grid item xs={12}>
         { 
		   <Table>
        <TableHead>
          <TableRow>
		  <TableCell style={{width: "10%"}}>Tid</TableCell>
            <TableCell style={{width: "70%"}}>Topic</TableCell>
            <TableCell style={{width: "10%"}}>Username</TableCell>
			<TableCell style={{width: "10%"}}>Operate</TableCell>
          </TableRow>
        </TableHead>
        <TableBody>
		
		{/*thread.username+","+thread.topic+","+currentusername+","+thread.tid*/}
		 { 
			 data.map((item,index) => (  
		          <TableRow id={'tr'+index} key={'tr'+index}>
				  <TableCell>
				  {item.split(',')[item.split(',').length-1]}{/*Extract the last element of the comma-separated array in the entire string*/}
				  </TableCell>
				  <TableCell>
				  <Link 
href={`/thread?topic=${item.split(",").slice(1, 2).join(",")}&username=${item.slice(0, item.indexOf(","))}&tid=${item.split(',')[item.split(',').length-1]}`}>
				  {item.split(",").slice(1, 2).join(",")}{/*Extract the second element of the comma-separated array in the entire string*/}
				  </Link>
				  </TableCell>
				  
		            <TableCell>{item.slice(0, item.indexOf(","))} {/*Extract the first element of the comma-separated array in the entire string*/}</TableCell>   
		            <TableCell>
					
					{/*Logic regarding the delete button*/}
					  { item.slice(0, item.indexOf(","))!=''&&item.slice(0, item.indexOf(","))===item.split(",").slice(-2)[0]  && 
					  <Link  onClick={() => myDelete(index)}>delete</Link>}  
					  <div style={{display: 'none'}} id={index}>
					  {item.split(',')[item.split(',').length-1]}  
 					  
					  </div>
					  </TableCell>   
					 </TableRow>
		        ))} 
				 
          
          {/* Add more rows as needed */}
        </TableBody>
      </Table>
	  }
          </Grid>
		  <Grid item xs={12}>
		    {}
		  </Grid>
          {}
        </Grid>
      </>
    </Container>
  );
};

export default Blog;
