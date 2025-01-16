import { hexToRGBA } from "@/utils/hex-to-rgba"

const listitem = theme => {
  return {
    MuiList: {
      styleOverrides: {
        root: {
          padding: 0,
        },
      },
    },
    MuiListSubheader: {
      styleOverrides: {
        root: {
          // backgroundColor: theme.palette.action.active,
        },
      },
    },
    MuiListItem: {
      styleOverrides: {
        root: {
          fontSize: "1rem",
        },
      },
    },
    MuiListItemText: {
      styleOverrides: {
        root: {
          fontSize: "1rem",
        },
      },
    },
    MuiListItemButton: {
      styleOverrides: {
        root: ({ ownerState }) => ({
          borderRadius: "1.25rem",
          padding: "0rem 1rem",

          "&:hover": {
            ...(!ownerState?.special && {
              background: `linear-gradient(90deg, ${hexToRGBA(theme.palette.primary.main, 0)} 0%, ${theme.palette.primary.main} 50%, ${hexToRGBA(theme.palette.primary.main, 0)} 100%)`,
              webkitBackgroundClip: "text",
              backgroundClip: "text",
              webkitTextFillColor: "transparent",
              color: "transparent",
            })
          },

          ...(ownerState?.active && {
            // gradient text color
            background: `linear-gradient(90deg, ${hexToRGBA(theme.palette.primary.main, 0)} 0%, ${theme.palette.primary.main} 50%, ${hexToRGBA(theme.palette.primary.main, 0)} 100%)`,
            webkitBackgroundClip: "text",
            backgroundClip: "text",
            webkitTextFillColor: "transparent",
            color: "transparent",
            display: "inline-block",

            "&::after": {
              content: '""',
              position: "absolute",
              bottom: "calc(-26%)",
              left: "calc(50% - 37.5%)",
              width: "75%",
              height: "0.05rem",
              // background: `linear-gradient(to right, ${theme.palette.primary.main}, ${theme.palette.info.main})`,
              background: `linear-gradient(90deg, ${hexToRGBA(theme.palette.primary.main, 0)} 0%, ${theme.palette.primary.main} 50%, ${hexToRGBA(theme.palette.primary.main, 0)} 100%)`,
            }
          }),

          ...(ownerState?.special && {
            // gradient text color
            border: "1px solid transparent",
            borderRadius: "0.5rem",
            background: `linear-gradient(${theme.palette.background.paper}, ${theme.palette.background.paper}), linear-gradient(to right, ${theme.palette[ownerState.color || "primary"].main}, ${theme.palette.info.main})`,
            backgroundClip: "padding-box, border-box",
            backgroundOrigin: "padding-box, border-box",

            "&:hover": {
              color: theme.palette.primary.main,
            },
          }),
        }),
      },
    },
    MuiListItemIcon: {
      styleOverrides: {
        root: {
          color: "inherit",
        },
      },
    },
  }
}

export default listitem
