import { hexToRGBA } from "@/utils/hex-to-rgba"

const card = theme => {
  return {
    MuiCard: {
      styleOverrides: {
        root: ({ ownerState }) => ({
          position: "relative",
          backgroundColor: theme.palette.background.paper,
          borderRadius: "0rem",
          border: "none",

          ...(
            ownerState.variant == "flat"
              ? {
              }
              : ownerState.variant == "special"
                ? {
                  background: theme.palette.background.default,
                  "&::before": {
                    content: '""',
                    display: "block",
                    width: "auto",
                    height: "auto",
                    minHeight: "50%",
                    maxHeight: "50%",
                    aspectRatio: 1,
                    position: "absolute",
                    top: "0px",
                    left: "0px",
                    borderTop: `2px solid ${hexToRGBA(theme.palette[ownerState.color || "primary"]?.main, 0.3)}`,
                    borderLeft: `2px solid ${hexToRGBA(theme.palette[ownerState.color || "primary"]?.main, 0.3)}`,
                  },
                  "&::after": {
                    content: '""',
                    display: "block",
                    width: "auto",
                    height: "auto",
                    minHeight: "50%",
                    maxHeight: "50%",
                    aspectRatio: 1,
                    position: "absolute",
                    bottom: "0px",
                    right: "0px",
                    borderBottom: `2px solid ${hexToRGBA(theme.palette[ownerState.color || "primary"]?.main, 0.3)}`,
                    borderRight: `2px solid ${hexToRGBA(theme.palette[ownerState.color || "primary"]?.main, 0.3)}`
                  }
                }
                : ownerState.variant == "gradient"
                  ? {
                    borderRadius: "1rem",
                    background: `linear-gradient(to right top, ${theme.palette[ownerState.color || "primary"].main} 0%, ${theme.palette.background.default} 75%)`,
                    padding: "1px",
                    clipPath: "polygon(0 0, 100% 0, 100% 80%, 70% 80%, 70% 100%, 0 100%)",

                    "& .MuiCardContent-root": {
                      borderRadius: "1rem",
                      background: theme.palette.background.default,
                      clipPath: "polygon(0 0, 100% 0, 100% 80%, 70% 80%, 70% 100%, 0 100%)",
                    }
                  }
                  : {
                    borderRadius: "0.5rem",
                    border: `1px solid ${hexToRGBA(theme.palette.border.light, 1)}`,
                  }
          ),
        }),
      },
    },
    MuiCardHeader: {
      styleOverrides: {
        root: ({ ownerState }) => ({
          padding: "0.5rem 1rem",
          borderBottom: `0.5px solid ${hexToRGBA(theme.palette.border.light, 1)}`,
          "& .MuiCardHeader-title": {
            color: theme.palette[ownerState.color || "text"]?.[!ownerState.color || ownerState.color == "text" ? "primary" : "main"],
          },
        }),
      },
    },
    MuiCardContent: {
      styleOverrides: {
        root: {
          padding: "1rem",
          paddingBottom: "1rem !important",
        },
      },
    },
    MuiCardActions: {
      styleOverrides: {
        root: {
          padding: "0.5rem 1rem",
          borderTop: `0.5px dashed ${hexToRGBA(theme.palette.border.light, 1)}`,
        },
      },
    }
  }
}

export default card
