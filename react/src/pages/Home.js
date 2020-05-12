import React, { useState, useEffect } from "react";
import "bootstrap/dist/css/bootstrap.min.css";
import Table from "react-bootstrap/Table";
import { connect } from "react-redux";
import { siswaShow, addSiswa, deleteSiswa } from "../action/siswaAction";

const Home = (props) => {
  const [siswa, setSiswa] = useState("");
  const [nama, setNama] = useState("");
  const [nisn, setNisn] = useState("");
  const [pendidikan, setPendidikan] = useState("");

  useEffect(() => {
    props.siswaShow();
  }, []);
  // console.log(props);

  const handleChangeNama = (e) => {
    setNama(e.target.value);
  };

  const handleChangeNisn = (e) => {
    setNisn(e.target.value);
  };

  const handleChangePendidikan = (e) => {
    setPendidikan(e.target.value);
  };

  const handleAdd = (e) => {
    e.preventDefault();
    let newSiswa = {
      id: props.siswa[props.siswa.length - 1] + 1,
      nama: nama,
      nisn: nisn,
      pendidikan: pendidikan,
    };
    props.addSiswa(newSiswa);
  };

  const handleDelete = (id) => {
    props.deleteSiswa(id);
  };

  console.log(props);
  return (
    <div className="container">
      <h1>Form Data Siswa</h1>
      <form className="row mt-3 ">
        <div className="col">
          nama
          <input value={nama} onChange={handleChangeNama} />
        </div>
        <div className="col">
          nisn
          <input value={nisn} onChange={handleChangeNisn} />
        </div>
        <div className="col">
          pendidikan
          <input value={pendidikan} onChange={handleChangePendidikan} />
        </div>

        <button onClick={handleAdd}>Add Pendidikan</button>
      </form>

      <Table className="mt-5" striped bordered hover>
        <thead>
          <th>nama</th>
          <th>nisn</th>
          <th>pendidikan</th>
          <th>action</th>
        </thead>
        <tbody>
          {props.siswa.map((item, index = 1) => (
            <tr>
              <td>{index + 1}</td>
              <td>{item.nama}</td>
              <td>{item.nisn}</td>
              <td>{item.kelas}</td>
              <td>
                <button
                  onClick={() => {
                    handleDelete(item.id);
                  }}
                >
                  delete
                </button>
              </td>
            </tr>
          ))}
        </tbody>
      </Table>
    </div>
  );
};

const mapStateToProps = (props) => {
  return {
    siswa: props.data,
  };
};

const mapDispatchToProps = { siswaShow, addSiswa, deleteSiswa };

export default connect(mapStateToProps, mapDispatchToProps)(Home);
