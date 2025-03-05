import React from "react";
import {
  Box,
  Avatar,
  Typography,
  Button,
  LinearProgress,
  Grid,
} from "@mui/material";
import { theme } from "@/configs/theme";

const xpProgress = 50;

const Profile = () => {
  return (
    <Grid container spacing={6}>
      <Grid item xs={12} md={4}>
        <Box
          sx={{
            display: "flex",
            flexDirection: "column",
            alignItems: "center",
            width: "100%",
          }}
        >
          <Box
            sx={{
              display: "flex",
              alignItems: "center",
              justifyContent: "space-between",
              width: "100%",
            }}
          >
            <Avatar
              sx={{
                width: 80,
                height: 80,
                borderRadius: 0,
              }}
              src="/images/profil.png"
            />

            <Box
              sx={{
                display: "flex",
                flexDirection: "column",
                alignItems: "flex-end",
              }}
            >
              <Button
                sx={{
                  "&:after": {
                    border: "none",
                  },
                  "&:before": {
                    border: "none",
                  },
                  background: "#1D1D1E",
                  color: theme.palette.secondary.light,
                  border: "1px solid transparent",
                  borderImageSource:
                    "linear-gradient(to bottom, #39D98A, #1B3B6F)",
                  borderImageSlice: 1,
                  transition: "all 0.3s ease-in-out",
                  height: 40,
                  width: 120,
                }}
              >
                Edit Profile
              </Button>

              <Typography
                variant="h5"
                sx={{ color: "#fff", mt: 1, textAlign: "right" }}
              >
                0xRbcFu4k...YOm4n
              </Typography>
            </Box>
          </Box>

          <Box sx={{ width: "100%", mt: 2 }}>
            <LinearProgress
              variant="determinate"
              value={xpProgress}
              sx={{
                height: 16,
                borderRadius: 1,
                backgroundColor: "rgba(255, 255, 255, 0.1)",
                "& .MuiLinearProgress-bar": {
                  backgroundImage: "linear-gradient(to right, #35373F, #39D98A)",
                  borderRadius: 1,
                  transition: "all 0.3s ease-in-out",
                },
              }}
            />
            <Typography
              variant="body1"
              sx={{
                color: theme.palette.secondary.main,
                mt: 0.5,
                display: "block",
                textAlign: "right",
                width: "100%",
                fontWeight: "bold",
              }}
            >
              {xpProgress} / 100 XP
            </Typography>
          </Box>
        </Box>
      </Grid>

      <Grid item xs={12} md={8}>
        <Box
          sx={{
            display: "flex",
            justifyContent: "space-between",
            alignItems: "center",
            width: "100%",
            gap: 2,
          }}
        >
          {[1, 2, 3, 4].map((index) => (
            <Box
              key={index}
              sx={{
                display: "flex",
                alignItems: "center",
                height: 30,
                gap: 1,
                padding: 2,
                border: `1px solid ${theme.palette.secondary.dark}`,
                borderRadius: 1,
              }}
            >
              <Box
                sx={{
                  backgroundColor: theme.palette.background.paper,
                  padding: 1,
                  borderRadius: 1,
                }}
              >
                <Avatar
                  sx={{
                    width: 30,
                    height: 30,
                    borderRadius: 0,
                  }}
                  src="/images/profil.png"
                />
              </Box>
              <Box
                sx={{
                  display: "flex",
                  flexDirection: "column",
                  alignItems: "space-between",
                  gap: 0.5,
                  marginLeft: 1,
                }}
              >
                <Typography
                  variant="h6"
                  sx={{
                    color: theme.palette.secondary.main,
                    fontSize: "14px",
                  }}
                >
                  Quest Name
                </Typography>
                <Typography
                  variant="body1"
                  sx={{
                    color: "#fff",
                  }}
                >
                  0
                </Typography>
              </Box>
            </Box>
          ))}
        </Box>
      </Grid>

      <Grid item xs={12} sx={{ display: "flex", justifyContent: "center", mt: 4 }}>
        <Box
          sx={{
            position: "relative",
            width: 200,
            height: 200,
            borderRadius: "50%",
            backgroundColor: "rgba(255, 255, 255, 0.1)",
            display: "flex",
            justifyContent: "center",
            alignItems: "center",
            
          }}
        >
          {[1, 2, 3].map((index) => (
            <Box
              key={index}
              sx={{
                position: "absolute",
                width: `${100 + index * 15}px`,
                height: `${100 + index * 15}px`,
                borderRadius: "50%",
                border: `1px solid rgba(255, 255, 255, 0.${10 + index * 3})`,
                animation: "pulse 2s infinite alternate ease-in-out",
              }}
            />
          ))}

          <Typography
            variant="h6"
            sx={{ color: "#fff", fontWeight: "bold", textAlign: "center" }}
          >
            🚀
          </Typography>
        </Box>
      </Grid>
    </Grid>
  );
};

export default Profile;
