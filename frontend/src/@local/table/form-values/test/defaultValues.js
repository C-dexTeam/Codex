import * as yup from 'yup'

export const testValues = {
  input: "",
  output: "",
};

export const testFilters = {
  page: 1,
  limit: 10,
};

export const testSchema = yup.object().shape({
  input: yup.string().required("Input is required"),
  output: yup.string().required("Output is required"),
});

export const testEditSchema = yup.object().shape({
  input: yup.string().required("Input is required"),
  output: yup.string().required("Output is required"),
});
