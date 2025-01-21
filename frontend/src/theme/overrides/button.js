import { hexToRGBA } from "@/utils/hex-to-rgba"

const button = theme => {
  return {
    MuiButton: {
      styleOverrides: {
        root: ({ ownerState }) => ({
          color: `${theme.palette.text.primary}`,
          background: `${theme.palette.background.paper}`,
          // outline: `0.1em solid ${theme.palette[ownerState.color || "primary"].main}`,
          // outlineOffset: "-0.1em",
          padding: "0.5em 1em",
          borderRadius: "0.5em",
          textTransform: "none",
          position: "relative",
          fontSize: "1em",
          fontHeight: "1.5em",
          zIndex: 1,

          "&:hover": {
            background: `${hexToRGBA(theme.palette[ownerState.color || "primary"].main, 0.1)} !important`,
            color: `${theme.palette[ownerState.color || "primary"].contrastText}`,
          },

          ...(
            ownerState.variant === "outlined"
              ? { // outlined button
                background: "inherit !important",
                outline: `0.1em solid ${theme.palette[ownerState.color || "primary"].main}`,
                outlineOffset: "-0.1em",
              }
              : ownerState.variant === "gradient"
                ? { // gradient button
                  background: `linear-gradient(90deg, ${theme.palette[ownerState.color || "primary"].main}, ${theme.palette[ownerState.color || "primary"].dark}) !important`,
                  color: `${theme.palette[ownerState.color || "primary"].contrastText} !important`,
                  outline: "none",
                  webkitTransition: "all 0.45s cubic-bezier(0.86, 0, 0.07, 1)",
                  transition: "all 0.45s cubic-bezier(0.86, 0, 0.07, 1)",

                  "&:hover": {
                    webkitTextFillColor: `${theme.palette[ownerState.color || "primary"].contrastText}`,
                    textFillColor: `${theme.palette[ownerState.color || "primary"].contrastText}`,
                    webkitTransition: "all 0.45s cubic-bezier(0.86, 0, 0.07, 1)",
                    transition: "all 0.45s cubic-bezier(0.86, 0, 0.07, 1)",

                    "&::before": {
                      background: `${theme.palette.background.default} !important`,
                      transition: "none",
                    },
                    "&::after": {
                      background: `${theme.palette.background.default} !important`,
                      transition: "none",
                    },
                  }
                }
                : ownerState.variant === "contained"
                  ? { // contained button
                    borderRadius: "1.25rem 0 1.25rem 0",
                    background: `${theme.palette[ownerState.color || "primary"].main} !important`,
                    color: `${theme.palette[ownerState.color || "primary"].contrastText} !important`,
                    outline: "none",
                    webkitTransition: "all 0.45s cubic-bezier(0.86, 0, 0.07, 1)",
                    transition: "all 0.45s cubic-bezier(0.86, 0, 0.07, 1)",

                    "&:hover": {
                      webkitTextFillColor: `${theme.palette[ownerState.color || "primary"].contrastText}`,
                      textFillColor: `${theme.palette[ownerState.color || "primary"].contrastText}`,
                      webkitTransition: "all 0.45s cubic-bezier(0.86, 0, 0.07, 1)",
                      transition: "all 0.45s cubic-bezier(0.86, 0, 0.07, 1)",

                      "&::before": {
                        background: `${theme.palette.background.default} !important`,
                        transition: "none",
                      },
                      "&::after": {
                        background: `${theme.palette.background.default} !important`,
                        transition: "none",
                      },
                    }
                  }
                  : { // default button
                    borderRadius: "0em",

                    "&::before": {
                      content: '""',
                      display: "block",
                      width: "15%",
                      height: "calc(100% - 4px)",
                      position: "absolute",
                      left: "0px",
                      borderTop: `2px solid ${hexToRGBA(theme.palette[ownerState.color || "primary"]?.main, 1)}`,
                      borderBottom: `2px solid ${hexToRGBA(theme.palette[ownerState.color || "primary"]?.main, 1)}`,
                      borderLeft: `2px solid ${hexToRGBA(theme.palette[ownerState.color || "primary"]?.main, 1)}`
                    },
                    "&::after": {
                      content: '""',
                      display: "block",
                      width: "15%",
                      height: "calc(100% - 4px)",
                      position: "absolute",
                      right: "0px",
                      borderTop: `2px solid ${hexToRGBA(theme.palette[ownerState.color || "primary"]?.main, 1)}`,
                      borderBottom: `2px solid ${hexToRGBA(theme.palette[ownerState.color || "primary"]?.main, 1)}`,
                      borderRight: `2px solid ${hexToRGBA(theme.palette[ownerState.color || "primary"]?.main, 1)}`
                    },
                  }
          ),

          // borderRadius: "1.25rem",
          // background: `${theme.palette[ownerState.color || "primary"].main} !important`,
          // color: `${theme.palette[ownerState.color || "primary"].contrastText} !important`,
        }),
      },
    }
  }
}

export default button
