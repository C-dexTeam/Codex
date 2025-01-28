import ReactMarkdown from 'react-markdown';
import { Box } from '@mui/material';

const MdCompiler = ({ markdown }) => {
  


  return (
      <Box
      
      >
        <ReactMarkdown>{(markdown)}</ReactMarkdown> 
      </Box>
  );
};

export default MdCompiler;
