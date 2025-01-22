
const typography = theme => {
  return {
    MuiTypography: {
      styleOverrides: {
        root: ({ ownerState }) => ({
          fontFamily: 'Montserrat Alternates',
          color: ownerState?.color && theme.palette[ownerState?.color]?.[ownerState?.color == "text" ? "primary" : "main"],
          // color: theme.palette.text.primary,
        }),
      }
    }
  }
}

export default typography
