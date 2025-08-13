-- Government Vehicle Assets Seed Data
INSERT INTO
    vehicle_asset (
        license_plate,
        stnk_status,
        bpkb_number,
        bpkb_status,
        chassis_number,
        machine_number,
        type_name,
        wheels_id,
        model_id,
        color_id,
        fuel_id,
        owning_id,
        brand_id,
        cc_capacity,
        manufacture_year,
        tax_due_date,
        last_tax_payment_date,
        current_owner,
        company_id,
        stnk_photo_path,
        bpkb_photo_path,
        owner_id_photo_path,
        vehicle_photo_path,
        payment_billing_photo_path,
        recommendation_document_path,
        e_sign_status,
        notes,
        status,
        created_at,
        updated_at,
        created_by
    )
VALUES
    -- 1. Official Sedan - Toyota Camry 2.5L Hybrid
    (
        'RI 1',
        'AVAILABLE',
        'BPKB112233445',
        'AVAILABLE',
        'CH0011223344',
        'MN9988776655',
        'Camry 2.5L Hybrid',
        3,
        (
            SELECT id
            FROM vehicle_model
            WHERE
                name = 'Sedan'
            LIMIT 1
        ),
        (
            SELECT id
            FROM vehicle_color
            WHERE
                name = 'Hitam'
            LIMIT 1
        ),
        (
            SELECT id
            FROM vehicle_fuel
            WHERE
                type = 'Hybrid'
            LIMIT 1
        ),
        (
            SELECT id
            FROM vehicle_owning
            WHERE
                type = 'Perusahaan'
            LIMIT 1
        ),
        (
            SELECT id
            FROM vehicle_brand
            WHERE
                name = 'Toyota'
            LIMIT 1
        ),
        2487,
        2023,
        EXTRACT(
            EPOCH
            FROM NOW() + interval '1 year'
        ),
        EXTRACT(
            EPOCH
            FROM NOW()
        ),
        'Kementerian Keuangan',
        (
            SELECT id
            FROM utilization_company
            WHERE
                name = 'PT Logistik Indonesia'
            LIMIT 1
        ),
        '/gov/photos/stnk_1.jpg',
        '/gov/photos/bpkb_1.jpg',
        '/gov/photos/owner_id_1.jpg',
        '/gov/photos/vehicle_1.jpg',
        '/gov/docs/billing_1.pdf',
        '/gov/docs/recommendation_1.pdf',
        false,
        'Kendaraan dinas eselon I',
        'completed',
        EXTRACT(
            EPOCH
            FROM NOW()
        ),
        EXTRACT(
            EPOCH
            FROM NOW()
        ),
        1
    ),
    -- 2. Nissan Skyline R35 GT-R
    (
        'RI 2',
        'AVAILABLE',
        'BPKB223344556',
        'AVAILABLE',
        'CH1122334455',
        'MN0099887766',
        'Skyline R35 GT-R',
        3,
        (
            SELECT id
            FROM vehicle_model
            WHERE
                name = 'Sedan'
            LIMIT 1
        ),
        (
            SELECT id
            FROM vehicle_color
            WHERE
                name = 'Putih'
            LIMIT 1
        ),
        (
            SELECT id
            FROM vehicle_fuel
            WHERE
                type = 'Pertamax Turbo'
            LIMIT 1
        ),
        (
            SELECT id
            FROM vehicle_owning
            WHERE
                type = 'Perusahaan'
            LIMIT 1
        ),
        (
            SELECT id
            FROM vehicle_brand
            WHERE
                name = 'Nissan'
            LIMIT 1
        ),
        3799,
        2022,
        EXTRACT(
            EPOCH
            FROM NOW() + interval '1 year'
        ),
        EXTRACT(
            EPOCH
            FROM NOW()
        ),
        'POLRI',
        (
            SELECT id
            FROM utilization_company
            WHERE
                name = 'PT Transportasi Makmur'
            LIMIT 1
        ),
        '/gov/photos/stnk_2.jpg',
        '/gov/photos/bpkb_2.jpg',
        '/gov/photos/owner_id_2.jpg',
        '/gov/photos/vehicle_2.jpg',
        '/gov/docs/billing_2.pdf',
        '/gov/docs/recommendation_2.pdf',
        true,
        'Kendaraan operasional khusus',
        'completed',
        EXTRACT(
            EPOCH
            FROM NOW()
        ),
        EXTRACT(
            EPOCH
            FROM NOW()
        ),
        1
    ),
    -- 3. Toyota Fortuner VRZ 2.8L Diesel AT
    (
        'RI 3',
        'NOT_AVAILABLE',
        'BPKB334455667',
        'AVAILABLE',
        'CH2233445566',
        'MN1188776655',
        'Fortuner VRZ 2.8L Diesel AT',
        3,
        (
            SELECT id
            FROM vehicle_model
            WHERE
                name = 'SUV'
            LIMIT 1
        ),
        (
            SELECT id
            FROM vehicle_color
            WHERE
                name = 'Silver'
            LIMIT 1
        ),
        (
            SELECT id
            FROM vehicle_fuel
            WHERE
                type = 'Biosolar (B20)'
            LIMIT 1
        ),
        (
            SELECT id
            FROM vehicle_owning
            WHERE
                type = 'Perusahaan'
            LIMIT 1
        ),
        (
            SELECT id
            FROM vehicle_brand
            WHERE
                name = 'Toyota'
            LIMIT 1
        ),
        2755,
        2023,
        EXTRACT(
            EPOCH
            FROM NOW() + interval '1 year'
        ),
        EXTRACT(
            EPOCH
            FROM NOW()
        ),
        'Kemenhub',
        (
            SELECT id
            FROM utilization_company
            WHERE
                name = 'PT Angkasa Jaya'
            LIMIT 1
        ),
        '/gov/photos/stnk_3.jpg',
        '/gov/photos/bpkb_3.jpg',
        '/gov/photos/owner_id_3.jpg',
        '/gov/photos/vehicle_3.jpg',
        '/gov/docs/billing_3.pdf',
        '/gov/docs/recommendation_3.pdf',
        false,
        'Kendaraan dinas transportasi darat',
        'unsigned',
        EXTRACT(
            EPOCH
            FROM NOW()
        ),
        EXTRACT(
            EPOCH
            FROM NOW()
        ),
        1
    ),
    -- 4. Mitsubishi Pajero Sport Dakar 4x4
    (
        'RI 4',
        'AVAILABLE',
        'BPKB445566778',
        'NOT_AVAILABLE',
        'CH3344556677',
        'MN2299887766',
        'Pajero Sport Dakar 4x4',
        3,
        (
            SELECT id
            FROM vehicle_model
            WHERE
                name = 'SUV'
            LIMIT 1
        ),
        (
            SELECT id
            FROM vehicle_color
            WHERE
                name = 'Putih'
            LIMIT 1
        ),
        (
            SELECT id
            FROM vehicle_fuel
            WHERE
                type = 'Biosolar (B20)'
            LIMIT 1
        ),
        (
            SELECT id
            FROM vehicle_owning
            WHERE
                type = 'Perusahaan'
            LIMIT 1
        ),
        (
            SELECT id
            FROM vehicle_brand
            WHERE
                name = 'Mitsubishi'
            LIMIT 1
        ),
        2847,
        2023,
        EXTRACT(
            EPOCH
            FROM NOW() + interval '1 year'
        ),
        EXTRACT(
            EPOCH
            FROM NOW()
        ),
        'Kementerian PUPR',
        (
            SELECT id
            FROM utilization_company
            WHERE
                name = 'PT Mitra Trans'
            LIMIT 1
        ),
        '/gov/photos/stnk_4.jpg',
        '/gov/photos/bpkb_4.jpg',
        '/gov/photos/owner_id_4.jpg',
        '/gov/photos/vehicle_4.jpg',
        '/gov/docs/billing_4.pdf',
        '/gov/docs/recommendation_4.pdf',
        true,
        'Kendaraan dinas lapangan',
        'unverified',
        EXTRACT(
            EPOCH
            FROM NOW()
        ),
        EXTRACT(
            EPOCH
            FROM NOW()
        ),
        1
    ),
    -- 5. Isuzu Traga 1.5L Pickup (Complete)
    (
        'RI 5',
        'AVAILABLE',
        'BPKB556677889',
        'AVAILABLE',
        'CH4455667788',
        'MN3399887766',
        'Traga 1.5L Pickup',
        3,
        (
            SELECT id
            FROM vehicle_model
            WHERE
                name = 'Pickup'
            LIMIT 1
        ),
        (
            SELECT id
            FROM vehicle_color
            WHERE
                name = 'Hitam'
            LIMIT 1
        ),
        (
            SELECT id
            FROM vehicle_fuel
            WHERE
                type = 'Pertamax'
            LIMIT 1
        ),
        (
            SELECT id
            FROM vehicle_owning
            WHERE
                type = 'Perusahaan'
            LIMIT 1
        ),
        (
            SELECT id
            FROM vehicle_brand
            WHERE
                name = 'Isuzu'
            LIMIT 1
        ),
        1498,
        2022,
        EXTRACT(
            EPOCH
            FROM NOW() + interval '1 year'
        ),
        EXTRACT(
            EPOCH
            FROM NOW()
        ),
        'Dinas Pertanian',
        (
            SELECT id
            FROM utilization_company
            WHERE
                name = 'PT Sejahtera Abadi'
            LIMIT 1
        ),
        '/gov/photos/stnk_5.jpg',
        '/gov/photos/bpkb_5.jpg',
        '/gov/photos/owner_id_5.jpg',
        '/gov/photos/vehicle_5.jpg',
        '/gov/docs/billing_5.pdf',
        '/gov/docs/recommendation_5.pdf',
        false,
        'Kendaraan dinas pertanian',
        'uncompleted',
        EXTRACT(
            EPOCH
            FROM NOW()
        ),
        EXTRACT(
            EPOCH
            FROM NOW()
        ),
        1
    ),
    -- 6. Toyota Hiace Ambulance 2.7L (Complete)
    (
        'RI 6',
        'AVAILABLE',
        'BPKB667788990',
        'AVAILABLE',
        'CH5566778899',
        'MN4499887766',
        'Hiace Ambulance 2.7L',
        3,
        (
            SELECT id
            FROM vehicle_model
            WHERE
                name = 'Minivan'
            LIMIT 1
        ),
        (
            SELECT id
            FROM vehicle_color
            WHERE
                name = 'Putih'
            LIMIT 1
        ),
        (
            SELECT id
            FROM vehicle_fuel
            WHERE
                type = 'Pertamax'
            LIMIT 1
        ),
        (
            SELECT id
            FROM vehicle_owning
            WHERE
                type = 'Perusahaan'
            LIMIT 1
        ),
        (
            SELECT id
            FROM vehicle_brand
            WHERE
                name = 'Toyota'
            LIMIT 1
        ),
        2694,
        2022,
        EXTRACT(
            EPOCH
            FROM NOW() + interval '1 year'
        ),
        EXTRACT(
            EPOCH
            FROM NOW()
        ),
        'Dinas Kesehatan',
        (
            SELECT id
            FROM utilization_company
            WHERE
                name = 'PT Medika Trans'
            LIMIT 1
        ),
        '/gov/photos/stnk_6.jpg',
        '/gov/photos/bpkb_6.jpg',
        '/gov/photos/owner_id_6.jpg',
        '/gov/photos/vehicle_6.jpg',
        '/gov/docs/billing_6.pdf',
        '/gov/docs/recommendation_6.pdf',
        true,
        'Kendaraan ambulans dinas',
        'completed',
        EXTRACT(
            EPOCH
            FROM NOW()
        ),
        EXTRACT(
            EPOCH
            FROM NOW()
        ),
        1
    ),
    -- 7-10: Additional vehicles...
    -- 7. Mitsubishi Colt L300 Minibus
    (
        'RI 7',
        'AVAILABLE',
        'BPKB778899001',
        'AVAILABLE',
        'CH6677889900',
        'MN5599887766',
        'Colt L300 Minibus',
        3,
        (
            SELECT id
            FROM vehicle_model
            WHERE
                name = 'Minibus'
            LIMIT 1
        ),
        (
            SELECT id
            FROM vehicle_color
            WHERE
                name = 'Biru'
            LIMIT 1
        ),
        (
            SELECT id
            FROM vehicle_fuel
            WHERE
                type = 'Pertamax'
            LIMIT 1
        ),
        (
            SELECT id
            FROM vehicle_owning
            WHERE
                type = 'Perusahaan'
            LIMIT 1
        ),
        (
            SELECT id
            FROM vehicle_brand
            WHERE
                name = 'Mitsubishi'
            LIMIT 1
        ),
        2477,
        2021,
        EXTRACT(
            EPOCH
            FROM NOW() + interval '1 year'
        ),
        EXTRACT(
            EPOCH
            FROM NOW()
        ),
        'Dinas Sosial',
        (
            SELECT id
            FROM utilization_company
            WHERE
                name = 'PT Angkasa Jaya'
            LIMIT 1
        ),
        '/gov/photos/stnk_7.jpg',
        '/gov/photos/bpkb_7.jpg',
        '/gov/photos/owner_id_7.jpg',
        '/gov/photos/vehicle_7.jpg',
        '/gov/docs/billing_7.pdf',
        '/gov/docs/recommendation_7.pdf',
        false,
        'Kendaraan dinas sosial',
        'paid',
        EXTRACT(
            EPOCH
            FROM NOW()
        ),
        EXTRACT(
            EPOCH
            FROM NOW()
        ),
        1
    ),
    -- 8. Hino Dutro 130 HD 4x2 (Complete)
    (
        'RI 8',
        'AVAILABLE',
        'BPKB889900112',
        'AVAILABLE',
        'CH7788990011',
        'MN6699887766',
        'Dutro 130 HD 4x2',
        4,
        (
            SELECT id
            FROM vehicle_model
            WHERE
                name = 'Truck'
            LIMIT 1
        ),
        (
            SELECT id
            FROM vehicle_color
            WHERE
                name = 'Biru'
            LIMIT 1
        ),
        (
            SELECT id
            FROM vehicle_fuel
            WHERE
                type = 'Biosolar (B20)'
            LIMIT 1
        ),
        (
            SELECT id
            FROM vehicle_owning
            WHERE
                type = 'Perusahaan'
            LIMIT 1
        ),
        (
            SELECT id
            FROM vehicle_brand
            WHERE
                name = 'Hino'
            LIMIT 1
        ),
        4009,
        2020,
        EXTRACT(
            EPOCH
            FROM NOW() + interval '1 year'
        ),
        EXTRACT(
            EPOCH
            FROM NOW()
        ),
        'Dinas PUPR',
        (
            SELECT id
            FROM utilization_company
            WHERE
                name = 'PT Logistik Indonesia'
            LIMIT 1
        ),
        '/gov/photos/stnk_8.jpg',
        '/gov/photos/bpkb_8.jpg',
        '/gov/photos/owner_id_8.jpg',
        '/gov/photos/vehicle_8.jpg',
        '/gov/docs/billing_8.pdf',
        '/gov/docs/recommendation_8.pdf',
        false,
        'Kendaraan logistik dinas',
        'paid',
        EXTRACT(
            EPOCH
            FROM NOW()
        ),
        EXTRACT(
            EPOCH
            FROM NOW()
        ),
        1
    ),
    -- 9. Yamaha XMAX 300 (Police Bike)
    (
        'RI 9',
        'AVAILABLE',
        'BPKB991122334',
        'AVAILABLE',
        'CH8899001122',
        'MN7799887766',
        'XMAX 300',
        1,
        (
            SELECT id
            FROM vehicle_model
            WHERE
                name = 'Skuter'
            LIMIT 1
        ),
        (
            SELECT id
            FROM vehicle_color
            WHERE
                name = 'Hitam'
            LIMIT 1
        ),
        (
            SELECT id
            FROM vehicle_fuel
            WHERE
                type = 'Pertamax'
            LIMIT 1
        ),
        (
            SELECT id
            FROM vehicle_owning
            WHERE
                type = 'Perusahaan'
            LIMIT 1
        ),
        (
            SELECT id
            FROM vehicle_brand
            WHERE
                name = 'Yamaha'
            LIMIT 1
        ),
        292,
        2023,
        EXTRACT(
            EPOCH
            FROM NOW() + interval '1 year'
        ),
        EXTRACT(
            EPOCH
            FROM NOW()
        ),
        'POLRI',
        (
            SELECT id
            FROM utilization_company
            WHERE
                name = 'PT Transportasi Makmur'
            LIMIT 1
        ),
        '/gov/photos/stnk_9.jpg',
        '/gov/photos/bpkb_9.jpg',
        '/gov/photos/owner_id_9.jpg',
        '/gov/photos/vehicle_9.jpg',
        '/gov/docs/billing_9.pdf',
        '/gov/docs/recommendation_9.pdf',
        true,
        'Sepeda motor patroli',
        'completed',
        EXTRACT(
            EPOCH
            FROM NOW()
        ),
        EXTRACT(
            EPOCH
            FROM NOW()
        ),
        1
    ),
    -- 10. Honda Beat 110 (Delivery Scooter)
    (
        'AD 1100 BB',
        'AVAILABLE',
        'BPKB001122334',
        'AVAILABLE',
        'CH9900112233',
        'MN8899776655',
        'Beat 110',
        1,
        (
            SELECT id
            FROM vehicle_model
            WHERE
                name = 'Skuter'
            LIMIT 1
        ),
        (
            SELECT id
            FROM vehicle_color
            WHERE
                name = 'Merah'
            LIMIT 1
        ),
        (
            SELECT id
            FROM vehicle_fuel
            WHERE
                type = 'Pertalite'
            LIMIT 1
        ),
        (
            SELECT id
            FROM vehicle_owning
            WHERE
                type = 'Perusahaan'
            LIMIT 1
        ),
        (
            SELECT id
            FROM vehicle_brand
            WHERE
                name = 'Honda'
            LIMIT 1
        ),
        110,
        2023,
        EXTRACT(
            EPOCH
            FROM NOW() + interval '1 year'
        ),
        EXTRACT(
            EPOCH
            FROM NOW()
        ),
        'Dinas Pos',
        (
            SELECT id
            FROM utilization_company
            WHERE
                name = 'PT Angkasa Jaya'
            LIMIT 1
        ),
        '/gov/photos/stnk_10.jpg',
        '/gov/photos/bpkb_10.jpg',
        '/gov/photos/owner_id_10.jpg',
        '/gov/photos/vehicle_10.jpg',
        '/gov/docs/billing_10.pdf',
        '/gov/docs/recommendation_10.pdf',
        false,
        'Kendaraan dinas pengiriman',
        'paid',
        EXTRACT(
            EPOCH
            FROM NOW()
        ),
        EXTRACT(
            EPOCH
            FROM NOW()
        ),
        1
    );