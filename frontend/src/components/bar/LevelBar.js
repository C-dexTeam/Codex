import { Box, Typography, useTheme } from "@mui/material"

const LevelBar = ({ level = 0, proggress = 0 }) => {

    const theme = useTheme()

    return (
        <Box
            sx={{
                position: "absolute",
                bottom: 0,
                width: "100%",
                height: "1rem",
                borderRadius: "0.25rem",
                backgroundColor: "background.default",
                border: theme => "0.5px solid " + theme.palette.border.secondary,
            }}
        >
            <Box
                sx={{
                    position: "absolute",
                    bottom: 0,
                    left: 0,
                    width: `${proggress}%`,
                    height: "1rem",
                    borderRadius: "0.25rem",
                    background: theme => `linear-gradient(90deg, ${theme.palette.background.paper} 0%, ${theme.palette.primary.main} 100%)`,
                }}
            ></Box>

            <Box
                sx={{
                    position: "absolute",
                    bottom: -6,
                    left: 4,
                }}
            >
                <svg xmlns="http://www.w3.org/2000/svg" width="56" height="22" viewBox="0 0 56 22" fill="none">
                    <g filter="url(#filter0_b_242_4224)">
                        <path d="M3.9041 0.647151C4.05095 0.257748 4.42361 0 4.83978 0H29.037H51.1602C51.5764 0 51.949 0.257749 52.0959 0.647151L55.8669 10.6472C55.9527 10.8746 55.9527 11.1254 55.8669 11.3528L52.0959 21.3528C51.949 21.7423 51.5764 22 51.1602 22H29.037H4.83978C4.42361 22 4.05095 21.7423 3.9041 21.3528L0.133061 11.3528C0.0473013 11.1254 0.0473012 10.8746 0.133061 10.6472L3.9041 0.647151Z" fill={theme.palette.background.paper} fillOpacity={0.25} />
                        <path d="M4.83978 0.5H29.037H51.1602C51.3683 0.5 51.5546 0.628875 51.6281 0.823575L55.3991 10.8236C55.442 10.9373 55.442 11.0627 55.3991 11.1764L51.6281 21.1764C51.5546 21.3711 51.3683 21.5 51.1602 21.5H29.037H4.83978C4.6317 21.5 4.44537 21.3711 4.37194 21.1764L0.600901 11.1764C0.558021 11.0627 0.558021 10.9373 0.600901 10.8236L4.37194 0.823575C4.44537 0.628874 4.6317 0.5 4.83978 0.5Z" stroke="url(#paint0_linear_242_4224)" />
                    </g>
                    <defs>
                        <filter id="filter0_b_242_4224" x="-1.93127" y="-2" width="59.8625" height="26" filterUnits="userSpaceOnUse" color-interpolation-filters="sRGB">
                            <feFlood flood-opacity="0" result="BackgroundImageFix" />
                            <feGaussianBlur in="BackgroundImageFix" stdDeviation="1" />
                            <feComposite in2="SourceAlpha" operator="in" result="effect1_backgroundBlur_242_4224" />
                            <feBlend mode="normal" in="SourceGraphic" in2="effect1_backgroundBlur_242_4224" result="shape" />
                        </filter>
                        <linearGradient id="paint0_linear_242_4224" x1="28" y1="-4.68577e-07" x2="28" y2="22" gradientUnits="userSpaceOnUse">
                            <stop stop-color="#3961D9" />
                            <stop offset="1" stop-color="#008F8C" />
                        </linearGradient>
                    </defs>
                </svg>
            </Box>

            <Box
                sx={{
                    position: "absolute",
                    bottom: 0,
                    left: 10 + (3 - level.toString().length) * 4,
                }}
            >
                <Typography variant="caption1">
                    lvl {level}
                </Typography>
            </Box>
        </Box>
    )
}

export default LevelBar