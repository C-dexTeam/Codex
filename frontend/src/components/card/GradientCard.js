import { Box, Button, Card } from '@mui/material'
import { useEffect, useRef, useState } from 'react'

const GradientCard = (props) => {

    const [btnWidth, setBtnWidth] = useState(null)
    const [btnHeight, setBtnHeight] = useState(null)

    const btnRef = useRef(null)

    useEffect(() => {
        if (btnRef.current) {
            console.log(btnRef.current.clientWidth);

            setBtnWidth(btnRef.current.clientWidth || null)
            setBtnHeight(btnRef.current.clientHeight || null)
        }
    }
    ), [btnRef.current]

    return (
        <Box sx={{ position: "relative" }}>
            <Card {...props} variant='gradient' btnWidth={btnWidth} btnHeight={btnHeight}>
                {props.children}

                <Box
                    sx={{ // If you change there values, you must change the same values in frontend/src/theme/overrides/card.js
                        width: `calc(100% - ${btnWidth || 96}px - 4px)`,
                        height: btnHeight + 4,
                        p: "0 1rem",
                        borderRadius: "0 0 0 calc(1rem - 1px)",
                    }}
                >
                    {props.description}
                </Box>
            </Card>

            <Button
                ref={btnRef}
                variant='contained'
                color='primary'
                sx={{
                    position: "absolute",
                    right: "0px",
                    bottom: "0px",
                }}
            >
                sadasdasd

            </Button>
        </Box>
    )
}

export default GradientCard