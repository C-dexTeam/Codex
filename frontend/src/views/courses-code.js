import React, { useState } from "react";
import {
  Container,
  Grid,
  Paper,
  Typography,
  Button,
  Card,
  Box,
} from "@mui/material";
import Editor from "@monaco-editor/react";
import { theme } from "@/configs/theme";
import MdCompiler from "@/components/mc-compiler";

function Code() {
  const [code, setCode] = useState(
    `pragma solidity >=0.5.0 <0.6.0;\n\ncontract HelloWorld {\n}`
  );

  const mdExample = `
    # Hello World
    This is a simple contract that prints "Hello World" to the console.
    \`\`\`solidity
    pragma solidity >=0.5.0 <0.6.0;
    contract HelloWorld {
      function hello() public pure returns (string memory) {
        return "Hello World";
      }
    }
    \`\`\`
    `;

  return (
    <Box maxWidth="lg">
      <Grid container spacing={3}>
        <Grid item xs={12} md={5}>
          <Box
            sx={{
              border: `2px solid ${theme.palette.secondary.light} `,
              borderRadius: "16px",
              padding: "16px",
            }}
          >
            <Box
              sx={{
                display: "flex",
                justifyContent: "start",
                alignItems: "center",
                borderRadius: "5px",
                gap: "16px",
              }}
            >
              <Typography
                variant="h6"
                sx={{
                  padding: "6px",
                  border: `1px solid ${theme.palette.secondary.light} `,
                  borderRadius: "5px",
                }}
              >
                Description
              </Typography>
              <Typography
                variant="h6"
                sx={{
                  padding: "6px",
                  border: `1px solid ${theme.palette.secondary.light} `,
                  borderRadius: "5px",
                }}
              >
                Description
              </Typography>
              <Typography
                variant="h6"
                sx={{
                  padding: "6px",
                  border: `1px solid ${theme.palette.secondary.light} `,
                  borderRadius: "5px",
                }}
              >
                Description
              </Typography>
            </Box>

            <Box
              sx={{
                display: "flex",
                flexDirection: "column",
                gap: "16px",
                marginTop: "32px",
              }}
            >
            <MdCompiler markdown={mdExample} />
            </Box>
          </Box>
        </Grid>

        <Grid item xs={12} md={7}>
          <Typography variant="h6">Solidity Editor</Typography>
          <Editor
            height="300px"
            defaultLanguage="sol"
            defaultValue={code}
            theme="vs-dark"
            onChange={(value) => setCode(value)}
          />
          <Button variant="contained" sx={{ mt: 2 }} color="primary">
            Compile Solidity
          </Button>
        </Grid>
      </Grid>
    </Box>
  );
}

export default Code;
