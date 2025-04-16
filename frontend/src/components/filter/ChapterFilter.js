import { Box, Button, FormControl, InputLabel, MenuItem, Select, TextField, Typography } from '@mui/material'
import { useState } from 'react'

const ChapterFilter = ({ filters, setFilters, onClose }) => {
    const [localFilters, setLocalFilters] = useState(filters)

    const handleChange = (e) => {
        const { name, value } = e.target
        setLocalFilters(prev => ({
            ...prev,
            [name]: value
        }))
    }

    const handleApply = () => {
        setFilters(localFilters)
        onClose()
    }

    const handleReset = () => {
        setLocalFilters({
            page: 1,
            limit: 10,
            title: '',
            languageID: '',
            rewardID: '',
            grantsExperience: '',
            active: ''
        })
    }

    return (
        <Box sx={{ p: 2, minWidth: 300 }}>
            <Typography variant="h6" sx={{ mb: 2 }}>Filter Chapters</Typography>
            
            <TextField
                fullWidth
                label="Title"
                name="title"
                value={localFilters.title}
                onChange={handleChange}
                sx={{ mb: 2 }}
            />

            <FormControl fullWidth sx={{ mb: 2 }}>
                <InputLabel>Language</InputLabel>
                <Select
                    name="languageID"
                    value={localFilters.languageID}
                    onChange={handleChange}
                    label="Language"
                >
                    <MenuItem value="">All</MenuItem>
                    <MenuItem value="EN">English</MenuItem>
                    <MenuItem value="TR">Turkish</MenuItem>
                </Select>
            </FormControl>

            <FormControl fullWidth sx={{ mb: 2 }}>
                <InputLabel>Grants Experience</InputLabel>
                <Select
                    name="grantsExperience"
                    value={localFilters.grantsExperience}
                    onChange={handleChange}
                    label="Grants Experience"
                >
                    <MenuItem value="">All</MenuItem>
                    <MenuItem value="true">Yes</MenuItem>
                    <MenuItem value="false">No</MenuItem>
                </Select>
            </FormControl>

            <FormControl fullWidth sx={{ mb: 2 }}>
                <InputLabel>Active</InputLabel>
                <Select
                    name="active"
                    value={localFilters.active}
                    onChange={handleChange}
                    label="Active"
                >
                    <MenuItem value="">All</MenuItem>
                    <MenuItem value="true">Yes</MenuItem>
                    <MenuItem value="false">No</MenuItem>
                </Select>
            </FormControl>

            <Box sx={{ display: 'flex', justifyContent: 'space-between', mt: 2 }}>
                <Button onClick={handleReset}>
                    Reset
                </Button>
                <Button onClick={handleApply} variant="contained">
                    Apply
                </Button>
            </Box>
        </Box>
    )
}

export default ChapterFilter 