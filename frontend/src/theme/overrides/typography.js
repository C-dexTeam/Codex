
const typography = theme => {
  return {
    MuiTypography: {
      styleOverrides: {
        root: ({ ownerState }) => ({
          fontFamily: '"Montserrat Alternates", sans-serif',
          color: ownerState?.color ? theme.palette[ownerState?.color]?.[ownerState?.color == "text" ? "primary" : "main"] : "inherit",
          // color: theme.palette.text.primary,
        }),
      }
    }
  }
}

export default typography
