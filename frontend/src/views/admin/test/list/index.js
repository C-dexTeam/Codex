import { testColums } from "@/@local/table/columns/test";
import ClassicTable from "@/components/tables/ClassicTable";
import { fetchTest, getFilters, getTest } from "@/store/admin/test";
import { useRouter } from "next/router";
import React, { useEffect } from "react";
import { useDispatch, useSelector } from "react-redux";

const TetsList = () => {
  const dispatch = useDispatch();
  const router = useRouter();
  const { chapterId } = router.query;
  const totalCount = 10
    const filters = useSelector(getFilters)


  // ** Selectors
  const test = useSelector(getTest);

  useEffect(() => {
    console.log("router.query.id", router.query);
    if (router.query.chapterId) {
      dispatch(fetchTest());
    }
  }, [router.query.chapterId, ]);

  return (
    <ClassicTable
                header={{
                    title: 'Test',
                    btnText: "Create Test",
                    btnClick: () => { router.push(`/admin/test/${chapterId}/create`) },
                    totalCount: totalCount || 0,
                }}
                rows={test || []}
                columns={testColums}
                getRowId={(row) => row._id || row.id || Math.random().toString(36).substr(2, 9)}
                pagination={{
                    page: filters?.page,
                    pageCount: Math.ceil((totalCount || 0) / filters?.limit),
                    setPage: v => _setFilters({ ...filters, page: v }),
                }}
            />
  );
};

export default TetsList;
