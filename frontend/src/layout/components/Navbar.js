import { Box, Button, Container } from "@mui/material";
import themeConfig from "@/configs/themeConfig";
import NavigationList from "./navigation";

const Navbar = () => {
  return (
    <Box
      sx={{
        background: theme => theme.palette.background.default,
      }}
    >
      <Container maxWidth="lg">
        <Box
          sx={{
            display: 'flex',
            position: 'relative',
            justifyContent: 'space-between',
            alignItems: 'center',
            p: '0.5rem 0rem',
            width: 'calc(100%)',
            height: 'calc(80px)',
            maxHeight: 'calc(80px)',
          }}
        >
          <Box component="div" sx={{ borderRadius: '1.25rem 0rem 1.25rem 0rem', mb: '0.5rem', textAlign: 'center' }}>
            {themeConfig.templateName}
          </Box>

          <Box sx={{
            background: theme => theme.palette.background.paper,
            borderRadius: "2rem",
            border: theme => `1px solid ${theme.palette.border.secondary}`,
            p: '0.5rem 1rem',
          }}>
            <NavigationList />
          </Box>

          <Box sx={{ display: "flex", gap: "1rem" }}>
            <Button color="info" variant="outlined">En</Button>
            <Button color="primary">Start Course</Button>
          </Box>
        </Box>
      </Container>
    </Box>
  );
};

export default Navbar;
