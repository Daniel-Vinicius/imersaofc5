import { GetStaticPaths, GetStaticProps } from "next";
import Router, { useRouter } from "next/router";

import { Typography, Card, CardHeader, CardContent, Box } from "@mui/material";

import useSWR from "swr";
import axios from "axios";

interface Order {
  id: string;
  amount: number;
  credit_card_number: string;
  credit_card_name: string;
  status: string;
}

interface OrdersShowPageProps {
  orders: Order[];
};

const fetcher = (url: string) => {
  return axios.get<Order>(url).then(res => res.data);
}

const OrdersShowPage = (props: OrdersShowPageProps) => {
  const router = useRouter();
  const { id } = router.query;
  const { data, error } = useSWR<Order>(`${process.env.NEXT_PUBLIC_API_HOST}/orders/${id}`, fetcher, {
    onError: (error) => {
      if (error.response.status === 401 || error.response.status === 403) {
        Router.push('/login');
      }
    }
  });

  return (
    data ? (
      <div style={{ height: 400, width: "100%" }}>
        <Typography component="h1" variant="h4">
          Minhas ordens
        </Typography>

        <Card>
          <CardHeader
            title="Order"
            subheader={data.id}
            titleTypographyProps={{ align: "center" }}
            subheaderTypographyProps={{ align: "center" }}
            sx={{ backgroundColor: (theme) => theme.palette.grey[700] }}
          />

          <CardContent>
            <Box
              sx={{
                display: "flex",
                justifyContent: "center",
                alignItems: "baseline",
                mb: 2,
              }}>
              <Typography component="h2" variant="h3" color="text.primary">
                R$ {data.amount}
              </Typography>
            </Box>
            <ul style={{ listStyle: "none" }}>
              <Typography component="li" variant="subtitle1">
                {data.credit_card_number}
              </Typography>
              <Typography component="li" variant="subtitle1">
                {data.credit_card_name}
              </Typography>
            </ul>
          </CardContent>
        </Card>
      </div>
    ) : null
  );
};

export default OrdersShowPage;

export const getStaticPaths: GetStaticPaths = async (context) => {
  return {
    paths: [],
    fallback: 'blocking',
  }
}

export const getStaticProps: GetStaticProps = async (context) => {
  return {
    props: {},
    revalidate: 20
  };
};
