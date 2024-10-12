import { createApi } from "@reduxjs/toolkit/query/react";
import { axiosBaseQuery } from "@app/api/baseQuery";
import { identityAPIHandler } from "@app/api/handlers";

const projectAPI = createApi({
    reducerPath: "projectAPI",
    baseQuery: axiosBaseQuery(identityAPIHandler),
    tagTypes: ["projects"],
    endpoints: (builder) => ({
        getAllProjects: builder.query<any, void>({
            query: () => ({
                url: "projects/all",
                method: "get",
            }),
            providesTags: ["projects"],
        }),
    }),
});

export const { useGetAllProjectsQuery } = projectAPI;

export default projectAPI;