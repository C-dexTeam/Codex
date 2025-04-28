import { Box, Button, Grid } from "@mui/material";
import DefaultTextField from "../components/DefaultTextField";
import { useEffect, useState } from "react";
import { validate } from "@/utils/validation";
import { showToast } from "@/utils/showToast";
import { useDispatch, useSelector } from "react-redux";
import { useRouter } from "next/router";
import {
  testEditSchema,
  testSchema,
} from "@/@local/table/form-values/test/defaultValues";
import { getCurrentTest, getLoading } from "@/store/admin/test";

const TestForm = ({
  values,
  setValues,
  isEdit = false,
  handleSubmit: _handleSubmit,
  errors: propErrors,
}) => {
  const [localErrors, setLocalErrors] = useState(null);
  const [isSubmitted, setIsSubmitted] = useState(false);

  const dispatch = useDispatch();
  const router = useRouter();

  const loading = useSelector(getLoading);
  const test = useSelector(getCurrentTest);

  useEffect(() => {
    if (values)
      validate(
        isEdit ? testEditSchema : testSchema,
        values,
        setIsSubmitted,
        setLocalErrors
      );
  }, [values]);

  const getError = (field) => {
    if (propErrors?.length) {
      const error = propErrors.find((err) => err.field === field);
      return error?.error;
    }
    return isSubmitted && localErrors?.[field] ? localErrors[field] : undefined;
  };

  const handleSubmit = () => {
    setIsSubmitted(true);

    if (localErrors && Object.keys(localErrors)?.length) {
      showToast("dismiss");
      showToast("error", "Please check the required fields.");
      return;
    }

    _handleSubmit({
      ...values,
      test: test?.id,
    });
  };

  return (
    <Grid container spacing={0}>
      <Grid item xs={12}>
        <DefaultTextField
          fullWidth
          type="text"
          name="input"
          label="Input Value"
          value={values?.input || ""}
          onChange={(e) =>
            setValues({
              ...values,
              input: e.target.value,
            })
          }
          required
          error={getError("input")}
        />
      </Grid>

      <Grid item xs={12}>
        <DefaultTextField
          fullWidth
          type="text"
          name="output"
          label="Output Value"
          value={values?.output || ""}
          onChange={(e) =>
            setValues({
              ...values,
              output: e.target.value,
            })
          }
          required
          error={getError("output")}
        />
      </Grid>

      <Grid item xs={12}>
        <Box sx={{ textAlign: "end" }}>
          <Button variant="outlined" onClick={handleSubmit} disabled={loading}>
            {isEdit ? "Update" : "Create"}
          </Button>
        </Box>
      </Grid>
    </Grid>
  );
};

export default TestForm;
