import React, {useState} from 'react'
import BasePage from "./BasePage";

function CertificatePage() {
    const [results, setResults] = useState(25)

    return (
        <BasePage
            isList={true}
            breadcrumbs={[
                {
                    name: "Dashboard",
                    to: "/"
                },
                {
                    name: "Certificates"
                }
            ]}
            title="Certificates"
            subtitle="A certificate object represents a public certificate/private key pair for an SSL certificate.
            These objects are used by Kong to handle SSL/TLS termination for encrypted requests. Certificates are
            optionally associated with SNI objects to tie a cert/key pair to one or more hostnames."
            results={results}
            setResults={setResults}
        >
            <div>CertificatePage</div>
        </BasePage>
    )
}

export default CertificatePage
