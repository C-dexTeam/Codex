
const grid = theme => {
  return {
    MuiGrid: {
      styleOverrides: {
        root: {
          marginLeft: "0rem",
          width: "calc(100% - 1rem)",
          "&>.MuiGrid-item": {
            paddingTop: "1rem",
            paddingLeft: "1rem",
          }
        },
      },
    }
  }
}

export default grid
