
const select = theme => {
  return {
    MuiSelect: {
      styleOverrides: {
        root: {
          width: '100%',

          "& fieldset": {
            border: `1px solid ${theme.palette.border.secondary}`,
            borderRadius: "0.5rem",
            padding: 0,

            "&:hover": {
              borderColor: theme.palette.border.main,
            },
          },

          "& .MuiInputBase-input": {
            padding: "0.5rem",
            zIndex: 9,
          },

          "& .MuiSelect-icon": {
            color: theme.palette.text.primary
          }
        },
      },
    },

    MuiMenuItem: {
      styleOverrides: {
        root: {
          width: "100%",

          "&:hover": {
          },

          "&.Mui-disabled": {
            background: theme.palette.action.disabledBackground,
            color: theme.palette.action.disabled
          },

          "&.Mui-focusVisible": {
          }
        }
      }
    }
  }
}

export default select