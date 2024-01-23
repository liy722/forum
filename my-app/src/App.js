import { BrowserRouter as Router, Route, Routes } from 'react-router-dom';
import './App.css';
import SignInSide from './SignInSide';
import Signup from './signup';
import Blog from './Blog';
import Block from './Block';
import Reply from './Reply';
import Thread from './Thread';
function App() {

  return(
   <Router>
        <Routes>
          <Route path="/" element={<SignInSide />} />
          <Route path="/signup" element={<Signup />} />
		  <Route path="/blog" element={<Blog />} />
		  <Route path="/block" element={<Block />} />
		  <Route path="/reply" element={<Reply />} />
		  <Route path="/thread" element={<Thread />} />
        </Routes>
      </Router>
  );
}

export default App;
