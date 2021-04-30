import Head from "next/head"

type MetaProps = {
  title: string
}

const Metadata = ({ title }: MetaProps): JSX.Element => {
  return (
    <>
      <Head>
        <title>{title}</title>
        <link rel="icon" href="/favicon.ico" />
        <meta charSet="utf-8" />
        <meta name="viewport" content="width=device-width, initial-scale=1" />
      </Head>
    </>
  )
}

export default Metadata
