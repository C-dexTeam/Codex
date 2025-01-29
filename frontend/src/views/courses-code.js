import React, { useState } from "react";
import { Grid, Typography, Button, Box, Divider } from "@mui/material";
import Editor from "@monaco-editor/react";
import { theme } from "@/configs/theme";
import Description from "@/components/code-change/description/Description"; 
import Template from "@/components/code-change/template/Template";
import HintDialog from "@/components/code-change/hint/Hint";

function Code() {
  const codeExample = `pragma solidity >=0.5.0 <0.6.0;\n\ncontract HelloWorld {\n}`;

  const markdownContent = `
   ## Chapter 2: Contracts  
Solidity's code is encapsulated in **contracts**.  
A contract is the fundamental building block of Ethereum applications — all variables and functions belong to a contract, and this will be the starting point of all your projects.  

An empty contract named **HelloWorld** would look like this:  

\`\`\`solidity
pragma solidity >=0.5.0 <0.6.0;

contract HelloWorld {
}
\`\`\`
`;

  const [code, setCode] = useState(codeExample);
  const [isCorrect, setIsCorrect] = useState(true);
  const [activeTab, setActiveTab] = useState("description");

  const [hintOpen, setHintOpen] = useState(false);
  const hintText = "No hint available at the moment. Try again later 🌌"
  ;



  return (
    <Box maxWidth="lg">
      <Grid container spacing={3}>
        <Grid item xs={12} md={5}>
          <Box
            sx={{
              border: `2px solid ${theme.palette.secondary.light}`,
              borderRadius: "16px",
              padding: "16px",
              height: "80vh",
              overflow: "auto",
              display: "flex",
              flexDirection: "column",
            }}
          >
            <Box
              sx={{
                display: "flex",
                justifyContent: "start",
                alignItems: "center",
                gap: "16px",
              }}
            >
              {[
                { text: "Description", icon: "📌", value: "description" },
                { text: "Examples", icon: "🚀", value: "expected" },
              ].map((item, index) => (
                <Button
                  key={index}
                  startIcon={<span>{item.icon}</span>}
                  variant={"outlined"}
                  onClick={() => setActiveTab(item.value)}
                  sx={{
                    color:
                      activeTab === item.value
                        ? theme.palette.success.main
                        : theme.palette.secondary.light,
                    fontSize: "13px",
                    outline: "none",
                    border: "1px solid",
                    padding: "0em 1.5em",
                  }}
                >
                  {item.text}
                </Button>
              ))}
              <Button 
                startIcon={<span>💡</span>}
                variant={"outlined"}
                onClick={() => setHintOpen(true)}
                sx={{
                  color: theme.palette.secondary.light,
                  fontSize: "13px",
                  outline: "none",
                  border: "1px solid",
                  padding: "0em 1.5em",
                }}
              >
                Hint
              </Button>

            </Box>

            <Box sx={{ marginTop: "24px", flexGrow: 1 }}>
              {activeTab === "description" ? (
                <Description markdownContent={markdownContent} />
              ) : (
                <Template />
              )}
            </Box>
          </Box>
        </Grid>

       
        <Grid item xs={12} md={7}>
          <Box
            sx={{
              border: `2px solid ${theme.palette.secondary.light}`,
              borderRadius: "16px",
              padding: "16px",
              height: "80vh",
              overflow: "auto",
              display: "flex",
              flexDirection: "column",
            }}
          >
            <Box
              sx={{
                display: "flex",
                justifyContent: "space-between",
                alignItems: "center",
                color: "white",
              }}
            >
              <Box sx={{ display: "flex", gap: "6px" }}>
                {["red", "yellow", "green"].map((color, index) => (
                  <Box
                    key={index}
                    sx={{
                      width: "20px",
                      height: "20px",
                      borderRadius: "50%",
                      backgroundColor: color,
                    }}
                  />
                ))}
              </Box>

              <Typography
                variant="h6"
                sx={{ color: `${theme.palette.primary.main}` }}
              >
                deneme.js
              </Typography>
            </Box>
            <Box
              sx={{
                marginTop: "16px",
              }}
            >
              <Divider />
            </Box>

            <Box
              sx={{
                marginTop: "16px",
                flexGrow: 1,
                borderRadius: "16px",
                overflow: "hidden",
              }}
            >
              <Editor
                height="100%"
                defaultLanguage="solidity"
                language="solidity"
                value={code}
                onChange={(value) => setCode(value)}
                theme="vs-dark"
                options={{
                  minimap: { enabled: false },
                  fontSize: 14,
                  padding: { top: 10 },
                  formatOnType: true,
                  formatOnPaste: true,
                  formatOnSave: true,
                }}
              />
            </Box>

            <Box
              sx={{
                marginTop: "16px",
              }}
            >
              <Divider />
            </Box>

            <Box
              sx={{
                display: "flex",
                alignItems: "center",
                justifyContent: "space-between",
                gap: "8px",
                marginTop: "16px",
              }}
            >
              <Box
                sx={{
                  flex: 1,
                  border: `2px solid ${
                    isCorrect
                      ? theme.palette.success.main
                      : theme.palette.error.main
                  }`,
                  borderRadius: "8px",
                  height: "100%",
                  backgroundColor: isCorrect
                    ? "rgba(0, 255, 0, 0.1)"
                    : "rgba(255, 0, 0, 0.1)",
                  display: "flex",
                  flexDirection: "column",
                  justifyContent: "center",
                }}
              >
                <Typography
                  variant="h5"
                  sx={{
                    color: isCorrect
                      ? theme.palette.success.main
                      : theme.palette.error.main,
                    margin: "16px",
                    display: "flex",
                  }}
                >
                  {isCorrect
                    ? "Congratulations, your solution is correct! You can proceed to other transactions!"
                    : "Error: There is an issue with your code. Please try again."}
                </Typography>
              </Box>

              <Box
                sx={{
                  display: "flex",
                  flexDirection: "column",
                  gap: "8px",
                }}
              >
                <Button variant="text" color="primary">
                  Run
                </Button>
                <Button variant="text" color="secondary">
                  Next
                </Button>
              </Box>
            </Box>
          </Box>
          <HintDialog open={hintOpen} onClose={() => setHintOpen(false)} hint={hintText} />
        </Grid>
      </Grid>
    </Box>
  );
}

export default Code;
