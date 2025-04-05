import { Box, Container, Divider, Grid, Typography } from "@mui/material";


const Footer = () => {
  return (
    <>
      <Box sx={{ background: theme => theme.palette.background.default }}>
        <Box sx={{ width: "100%", borderTop: theme => "0.5px solid " + theme.palette.border.secondary }}></Box>
        <Box sx={{ width: "100%", display: "flex", justifyContent: "center" }}>
          <Box sx={{ width: "50%" }}>
            <Divider color="info" />
          </Box>
        </Box>

        <Container>
          <Grid container sx={{ pt: 2, pb: 4 }} spacing={2.5}>
            <Grid item xs={12} md={8}>
              <Box sx={{ display: "flex", gap: "1rem", width: '100%' }}>
                <Typography variant="body1">© 2024 Codex Team</Typography>
                <Typography variant="body1">Privacy Policy</Typography>
                <Typography variant="body1">Terms of Use</Typography>
              </Box>

              <Box sx={{ mt: 2 }}>
                <Typography variant="caption1" sx={{ textAlign: "justify" }}>
                  Codex is very important project for solana. The codex team was here property of their respective owners. property of their respective owners. Unless otherwise noted, use of third party logos does not imply endorsement of, sponsorship of, or affiliation with Cobalt. Cobalt is a financial technology company, not a bank. Banking services are provided by Celtic Bank and Evolve Bank & Trust®, Members FDIC.
                </Typography>
              </Box>
            </Grid>
            <Grid item xs={12} md={4}>b</Grid>
          </Grid>
        </Container>
      </Box>
    </>
  );
};

export default Footer;
