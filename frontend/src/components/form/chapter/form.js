import { Button, Grid, TextField, FormControlLabel, Checkbox, MenuItem, Box, Typography } from '@mui/material'
import { useState, useEffect } from 'react'
import { useDispatch, useSelector } from 'react-redux'
import { fetchCourses, getCourses } from '@/store/admin/courses'
import { fetchRewards, getRewards } from '@/store/admin/rewards'
import { fetchLanguages, getLanguages } from '@/store/admin/languages'
import { validate } from '@/utils/validation'
import { showToast } from '@/utils/showToast'
import { chapterSchema } from '@/@local/table/form-values/chapter/defaultValues'
import DefaultTextField from '../components/DefaultTextField'
import DefaultSelect from '../components/DefaultSelect'

const ChapterForm = ({ values, setValues, handleSubmit: _handleSubmit, isEdit = false }) => {
    const [localErrors, setLocalErrors] = useState(null)
    const [isSubmitted, setIsSubmitted] = useState(false)
    const dispatch = useDispatch()
    const courses = useSelector(getCourses)
    const languages = useSelector(getLanguages)
    // const { rewards } = useSelector(getRewards)

    useEffect(() => {
        dispatch(fetchCourses())
        dispatch(fetchLanguages())
        // dispatch(fetchRewards())
    }, [])

    useEffect(() => {
        if (values) {
            validate(chapterSchema, values, setIsSubmitted, setLocalErrors)
        }
    }, [values])

    const handleSubmit = (e) => {
        e.preventDefault()
        setIsSubmitted(true)
        
        if (localErrors && Object.keys(localErrors)?.length) {
            showToast("dismiss")
            showToast("error", "Lütfen gerekli alanları kontrol edin.")
            return
        }

        _handleSubmit({
            ...values,
            languageID: languages?.find(language => language.value == "EN")?.id,
        })
    }

    const getError = (field) => {
        return isSubmitted && localErrors?.[field] ? localErrors[field] : undefined
    }

    const handleChange = (e) => {
        const { name, value, checked } = e.target
        setValues({
            ...values,
            [name]: e.target.type === 'checkbox' ? checked : value
        })
    }

    return (
        <Grid container spacing={3}>
            <Grid item xs={12}>
                <DefaultSelect
                    required
                    label="Course"
                    id='courseID'
                    firstSelect={"-- Select a course --"}
                    value={values?.courseID}
                    onChange={e => setValues({ ...values, courseID: e.target.value })}
                    items={
                        courses && courses?.length > 0 &&
                        courses?.map((item, index) => {
                            return (
                                <MenuItem key={item?.id} value={item?.id}>
                                    <Box sx={{ display: "flex", alignItems: "center", justifyContent: "space-between", gap: "1rem", width: "100%" }}>
                                        <Box sx={{ display: "flex", alignItems: "center", gap: "1rem" }}>
                                            {item?.imagePath && (
                                                <img src={"/api/" + item?.imagePath} style={{ width: 40, height: 40, objectFit: 'contain' }} />
                                            )}

                                            <Typography variant='body' component="span">
                                                {item?.title}
                                            </Typography>
                                        </Box>

                                        <Typography variant='caption' component="span">
                                            {item?.chapterCount} chapters
                                        </Typography>
                                    </Box>
                                </MenuItem>
                            )
                        })
                    }
                    error={getError('courseID')}
                />
            </Grid>

            {/* <Grid item xs={12}>
                    <DefaultSelect
                        required
                        label="Language"
                        id='languageID'
                        firstSelect={"-- Select a language --"}
                        value={values?.languageID}
                        onChange={e => setValues({ ...values, languageID: e.target.value })}
                        items={
                            languages && languages?.length > 0 &&
                            languages?.map((item, index) => {
                                return (
                                    <MenuItem key={item?.id} value={item?.id}>
                                        {item?.value}
                                    </MenuItem>
                                )
                            })
                        }
                        error={getError('languageID')}
                    />
                </Grid> */}

            <Grid item xs={12}>
                <DefaultTextField
                    fullWidth
                    type="text"
                    label="Title"
                    name="title"
                    value={values.title}
                    onChange={handleChange}
                    required
                    error={getError('title')}
                />
            </Grid>

            <Grid item xs={12}>
                <DefaultTextField
                    fullWidth
                    label="Description"
                    name="description"
                    value={values.description}
                    onChange={handleChange}
                    multiline
                    rows={4}
                    error={getError('description')}
                />
            </Grid>

            <Grid item xs={12}>
                <DefaultTextField
                    fullWidth
                    label="Content"
                    name="content"
                    value={values.content}
                    onChange={handleChange}
                    multiline
                    rows={6}
                    required
                    error={getError('content')}
                />
            </Grid>

            <Grid item xs={12}>
                <DefaultSelect
                    fullWidth
                    label="Reward"
                    name="rewardID"
                    value={values.rewardID}
                    onChange={handleChange}
                    firstSelect="-- Select a reward --"
                    // options={rewards}
                    error={getError('rewardID')}
                />
            </Grid>

            <Grid item xs={12}>
                <DefaultTextField
                    fullWidth
                    label="Function Name"
                    name="funcName"
                    value={values.funcName}
                    onChange={handleChange}
                    required
                    error={getError('funcName')}
                />
            </Grid>

            <Grid item xs={12}>
                <DefaultTextField
                    fullWidth
                    label="Frontend Template"
                    name="frontendTemplate"
                    value={values.frontendTemplate}
                    onChange={handleChange}
                    multiline
                    rows={4}
                    error={getError('frontendTemplate')}
                />
            </Grid>

            <Grid item xs={12}>
                <DefaultTextField
                    fullWidth
                    label="Docker Template"
                    name="dockerTemplate"
                    value={values.dockerTemplate}
                    onChange={handleChange}
                    multiline
                    rows={4}
                    error={getError('dockerTemplate')}
                />
            </Grid>

            <Grid item xs={12}>
                <DefaultTextField
                    fullWidth
                    label="Check Template"
                    name="checkTemplate"
                    value={values.checkTemplate}
                    onChange={handleChange}
                    multiline
                    rows={4}
                    required
                    error={getError('checkTemplate')}
                />
            </Grid>

            <Grid item xs={12}>
                <DefaultTextField
                    fullWidth
                    type="number"
                    label="Reward Amount"
                    name="rewardAmount"
                    value={values.rewardAmount}
                    onChange={handleChange}
                    error={getError('rewardAmount')}
                />
            </Grid>

            <Grid item xs={12}>
                <DefaultTextField
                    fullWidth
                    type="number"
                    label="Order"
                    name="order"
                    value={values.order}
                    onChange={handleChange}
                    required
                    error={getError('order')}
                />
            </Grid>

            <Grid item xs={12}>
                <FormControlLabel
                    control={
                        <Checkbox
                            name="grantsExperience"
                            checked={values.grantsExperience}
                            onChange={handleChange}
                        />
                    }
                    label="Grants Experience"
                />
            </Grid>

            <Grid item xs={12}>
                <FormControlLabel
                    control={
                        <Checkbox
                            name="active"
                            checked={values.active}
                            onChange={handleChange}
                        />
                    }
                    label="Active"
                />
            </Grid>

            <Grid item xs={12}>
                <Button
                    type="submit"
                    variant="contained"
                    color="primary"
                    onClick={handleSubmit}
                    fullWidth
                >
                    {isEdit ? 'Update Chapter' : 'Create Chapter'}
                </Button>
            </Grid>
        </Grid>
    )
}

export default ChapterForm