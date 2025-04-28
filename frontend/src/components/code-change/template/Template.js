import React, { useEffect, useState } from "react";
import {
  Box,
  Typography,
  Accordion,
  AccordionSummary,
  AccordionDetails,
  TextField,
  Grid,
} from "@mui/material";
import ExpandMoreIcon from "@mui/icons-material/ExpandMore";
import { theme } from "@/configs/theme";
import { useDispatch, useSelector } from "react-redux";
import { useRouter } from "next/router";
import { getTest } from "@/store/test/testSlice";

const TestCases = () => {
  const dispatch = useDispatch();
  const router = useRouter();
  const { id } = router.query;

  const { test: testSlice } = useSelector((state) => state);

  useEffect(() => {
    if (router.isReady) {
      dispatch(getTest({ id: router.query.code }));
    }
  }, [router.isReady, router.query.code]);

  const fallbackCases = [
    { id: 1, input: "I'm Sorry", output: "here is empty" },
  ];

  const testCases =
    testSlice?.data?.data && testSlice.data.data.length > 0
      ? testSlice.data.data
      : fallbackCases;

  return (
    <Box sx={{ height: "100%", overflow: "auto" }}>
      <Typography variant="h4" color={`${theme.palette.primary.main}`}>
        Expected Test Cases
      </Typography>
      {testCases?.map((testCase, index) => (
        <Accordion
          key={testCase.id}
          sx={{
            backgroundColor: `${theme.palette.action.hover}`,
            color: "white",
            borderRadius: "8px",
            mb: 1,
          }}
        >
          <AccordionSummary
            expandIcon={<ExpandMoreIcon sx={{ color: "white" }} />}
          >
            <Typography>Case {index + 1}</Typography>
          </AccordionSummary>
          <AccordionDetails>
            <Grid container spacing={2}>
              <Grid item xs={12}>
                <Typography
                  variant="body1"
                  color={`${theme.palette.secondary.main}`}
                >
                  Input
                </Typography>

                <TextField
                  fullWidth
                  variant="outlined"
                  value={testCase.input}
                  InputProps={{ readOnly: true }}
                  sx={{
                    backgroundColor: `${theme.palette.action.hover}`,
                    borderRadius: "8px",
                    "& .MuiOutlinedInput-root": {
                      color: "white",
                      borderColor: `${theme.palette.secondary.main}`,
                    },
                  }}
                />
              </Grid>

              <Grid item xs={12}>
                <Typography
                  variant="body1"
                  color={`${theme.palette.secondary.main}`}
                >
                  Output
                </Typography>
                <TextField
                  fullWidth
                  variant="outlined"
                  value={testCase.output}
                  InputProps={{ readOnly: true }}
                  sx={{
                    backgroundColor: `${theme.palette.action.hover}`,
                    borderRadius: "8px",
                    "& .MuiOutlinedInput-root": {
                      color: "white",
                      borderColor: `${theme.palette.secondary.main}`,
                    },
                  }}
                />
              </Grid>
            </Grid>
          </AccordionDetails>
        </Accordion>
      ))}
    </Box>
  );
};

export default TestCases;
