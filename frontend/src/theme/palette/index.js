import { hexToRGBA } from "@/utils/hex-to-rgba";

const text = {
    primary: "#f0f0f0",
    secondary: "#0F0F0F",
};

const border = {
    main: "#A1A6B4",
    light: "#DEE0E8",
    secondary: "#35373F"
}

const palette = {
    action: {
        active: hexToRGBA(border.main, 0.12),
        disabled: border.dark,
        disabledBackground: border.light,
        focus: border.light,
        hover: hexToRGBA(border.main, 0.24),
        selected: border.light,
    },
    common: {
        black: text.primary,
        white: text.secondary,
    },
    divider: border.main,
    primary: {
        main: "#39D98A",
        light: "#ABFFD8",
        dark: "#01572F",
        contrastText: text.secondary,
    },
    secondary: {
        main: "#A1A6B4",
        light: "#DEE0E8",
        dark: "#35373F",
        contrastText: text.primary,
    },
    success: {
        main: "#39D98A",
        light: "#ABFFD8",
        dark: "#01572F",
        contrastText: text.primary,
    },
    warning: {
        main: "#E3BE44",
        light: "#FFF3AF",
        dark: "#816B0E",
        contrastText: text.primary,
    },
    info: {
        main: "#1B3B6F",
        light: "#AABBD8",
        dark: "#00122F",
        contrastText: text.primary,
    },
    error: {
        main: "#D72638",
        light: "#FFAAAD",
        dark: "#570003",
        contrastText: text.primary,
    },
    background: {
        default: "#0B0C10",
        paper: "#1D1F2B",
    },
    border: {
        main: border.main,
        light: border.light,
        secondary: border.secondary,
    },
    text: {
        primary: text.primary,
        secondary: text.secondary,
        disabled: border.dark,
    },
};

export default palette;