import React, { useEffect, useState } from "react";
import { Grid, Typography, Button, Box, Divider } from "@mui/material";
import Editor from "@monaco-editor/react";
import { theme } from "@/configs/theme";
import Description from "@/components/code-change/description/Description";
import Template from "@/components/code-change/template/Template";
import HintDialog from "@/components/code-change/hint/Hint";
import { useRouter } from "next/router";
import { useDispatch, useSelector } from "react-redux";
import { getChaptersByID } from "@/store/chapters/chaptersSlice";

function Code() {
  const [code, setCode] = useState("");
  const [isCorrect, setIsCorrect] = useState(true);
  const [activeTab, setActiveTab] = useState("description");

  const [hintOpen, setHintOpen] = useState(false);
  const hintText = "No hint available at the moment. Try again later 🌌";
  const router = useRouter();
  const dispatch = useDispatch();
  const { chapters: chaptersSlice } = useSelector((state) => state);

  useEffect(() => {
    if (router.isReady) {
      dispatch(getChaptersByID({ id: router.query.code }));
    }
  }, [router.isReady, router.query.code]);

  console.log(chaptersSlice.data.data);

  const testsExist = chaptersSlice?.data?.data?.tests?.length > 0;

  const markdownContent =
    chaptersSlice?.data?.data?.content ||
    "🚀 Oops! This section seems **empty**. Let's add some useful content for now:\n\n" +
      "## Welcome!  \n" +
      "In this chapter, you will learn the basics of Solidity. Ready to write your first smart contract? Let's get started! 🎉\n\n" +
      "```solidity\n" +
      "pragma solidity >=0.5.0 <0.6.0;\n\n" +
      "contract HelloWorld {\n" +
      '    string public message = "Hello, Blockchain!";\n' +
      "}\n" +
      "```";

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
              <Button
                startIcon={<span>📌</span>}
                variant={"outlined"}
                onClick={() => setActiveTab("description")}
                sx={{
                  color:
                    activeTab === "description"
                      ? theme.palette.success.main
                      : theme.palette.secondary.light,
                  fontSize: "13px",
                  outline: "none",
                  border: "1px solid",
                  padding: "0em 1.5em",
                }}
              >
                Description
              </Button>

              {testsExist && (
                <Button
                  startIcon={<span>🚀</span>}
                  variant={"outlined"}
                  onClick={() => setActiveTab("expected")}
                  sx={{
                    color:
                      activeTab === "expected"
                        ? theme.palette.success.main
                        : theme.palette.secondary.light,
                    fontSize: "13px",
                    outline: "none",
                    border: "1px solid",
                    padding: "0em 1.5em",
                  }}
                >
                  Examples
                </Button>
              )}

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
                <Template /> //this code will updated when the test template added
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
                value={chaptersSlice?.data?.data?.frontendTemplate}
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
          <HintDialog
            open={hintOpen}
            onClose={() => setHintOpen(false)}
            hint={hintText}
          />
        </Grid>
      </Grid>
    </Box>
  );
}

export default Code;
