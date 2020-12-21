import "./App.css";

function PatientInfo(props) {
    console.log("THJIS HERE WAS RAN");

    {props.patientList.map((patient, index) => {
        return (
            <tr>
                <td>patient.name</td>
                <td>patient.dob</td>
                <td>patient.phone</td>
                <td>patient.email</td>
                <td>patient.address</td>
                <td>patient.time</td>
            </tr>
        )
    })
    }
}

export default PatientInfo;