import { Grid } from "@mui/material"
import FileUploaderSingle from "../components/FileUploaderSingle"
import DefaultTextField from "../components/DefaultTextField"
import { useEffect, useState } from "react"
import { validate } from "@/utils/validation"
import { showToast } from "@/utils/showToast"
import { courseSchema } from "@/@local/table/form-values/event/defaultValues"

const AddCourseForm = ({
    values,
    setValues,
    edit = false,
    handleSubmit: _handleSubmit,
}) => {
    const [errors, setErrors] = useState(null);
    const [isSubmitted, setIsSubmitted] = useState(false);
    const [files, setFiles] = useState([])

    const handleSubmit = () => {
        setIsSubmitted(true);
        if (errors && Object.keys(errors)?.length) {
            showToast("dismiss");
            showToast("error", "Lütfen gerekli alanları kontrol edin.");
            return;
        }

        _handleSubmit();
    };

    useEffect(() => {
        if (values) validate(courseSchema, values, setIsSubmitted, setErrors);
    }, [values]);

    return (
        <Grid container spacing={2}>
            <Grid item xs={12}>
                <FileUploaderSingle
                    files={files}
                    setFiles={setFiles}
                    text="Upload a cover image for the course"
                />
            </Grid>

            <Grid item xs={12}>
                <DefaultTextField
                fullWidth
                    type="text"
                    name="title"
                    label="Title"
                    value={values?.title}
                    onChange={(e) =>
                        setValues({
                            ...values,
                            title: e.target.value,
                        })
                    }
                    required
                    error={
                        isSubmitted && errors?.title
                            ? errors?.title
                            : undefined
                    }
                />
            </Grid>
        </Grid>
    )
}

export default AddCourseForm