use crate::{application::APP, cellars, cork, prelude::*};
use abscissa_core::{clap::Parser, Command, Runnable};
use deep_space::Address;
use ethers::abi::AbiEncode;
use steward_abi::cellar::*;

/// Fees Distributor subcommand
#[derive(Command, Debug, Parser)]
#[clap(
    long_about = "DESCRIPTION\n\nCalls setFeesDistributor() on the target cellar contract at the specified block height.\nFor more information see https://github.com/PeggyJV/cellar-contracts/blob/main/src/base/Cellar.sol"
)]
pub struct SetFeesDistributorCmd {
    #[clap(short = 'n', long)]
    /// Fee distributor's address
    new_fees_distributor: Address,

    /// Target cellar for scheduled cork.
    #[clap(short, long)]
    cellar_id: String,

    /// Block height to schedule cork.
    #[clap(short = 'b', long)]
    height: u64,
}

impl Runnable for SetFeesDistributorCmd {
    fn run(&self) {
        abscissa_tokio::run_with_actix(&APP, async {
            let mut address = self.new_fees_distributor.as_bytes().to_vec();

            while address.len() < 32 {
                address.insert(0, 0u8);
            }

            let mut address_slice: [u8; 32] = Default::default();
            address_slice.copy_from_slice(&address[..]);
            let call = SetFeesDistributorCall {
                new_fees_distributor: address_slice,
            };
            let encoded_call = CellarCalls::SetFeesDistributor(call).encode();

            cellars::validate_cellar_id(&self.cellar_id).unwrap_or_else(|err| {
                status_err!("invalid cellar ID: {}", err);
                std::process::exit(1);
            });

            cork::schedule_cork(self.cellar_id.clone(), encoded_call, self.height)
                .await
                .unwrap_or_else(|err| {
                    status_err!("error while submitting scheduled cork. please verify that the scheduling failed before attempting again: {}", err);
                    std::process::exit(1);
                })
        })
        .unwrap_or_else(|e| {
            status_err!("executor exited with error: {}", e);
            std::process::exit(1);
        });
    }
}
